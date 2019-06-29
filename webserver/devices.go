package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func NewDeviceController(g graph.Graph) DeviceController {
	if g == nil {
		panic("Graph should not be nil")
	}
	return DeviceController{
		graph: g,
	}
}

type DeviceController struct {
	graph graph.Graph
}

type DeviceResponse struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Type  string        `json:"type"`
	State devices.State `json:"state"`
}

type DetailedDeviceResponse struct {
	DeviceResponse
	DisplayOpts interface{} `json:"display_opts,omitempty"`
	InputPorts  []string    `json:"input_ports"`
	OutputPorts []string    `json:"output_ports"`
}

type DeviceUpdateParams struct {
	DisplayOpts interface{} `json:"display_opts,omitempty"`
}

// List all devices in the device graph
func (c DeviceController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	var devices []DeviceResponse

	// For each device in the device graph, get the corresponding DeviceResponse struct
	for _, device := range c.graph.All(ctx) {
		devices = append(devices, deviceToDeviceResponse(device))
	}

	utils.JSON(ctx, resp, devices)
	return nil
}

// Show a specific device of the device graph
func (c DeviceController) Show(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var dev models.Device
	err := document.Find(ctx, models.DeviceCollection, id, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to find device")
	}

	// Fetch the device in the device graph
	device, err := c.graph.Get(ctx, params["id"])
	if err != nil { // The device graph failed to find the device
		// If the device graph does not exist in the device graph
		if err == graph.ErrNotFound {
			utils.Err(ctx, resp, http.StatusNotFound, "not found")
			return nil
		} else {
			// It's another error => 500
			return errors.Wrap(err, "fail to find device")
		}
	}

	// Convert the graph device representation to a DetailedDeviceResponse
	deviceResp := DetailedDeviceResponse{
		DeviceResponse: deviceToDeviceResponse(device),
		InputPorts:     device.InputPorts(),
		OutputPorts:    device.OutputPorts(),
		DisplayOpts:    dev.DisplayOpts,
	}

	utils.JSON(ctx, resp, deviceResp)
	return nil
}

// Create and add a device to the device graph
func (c DeviceController) Create(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	device := models.Device{}
	log.Info("Start device import")

	// Decode the generic model
	err := json.NewDecoder(req.Body).Decode(&device)
	if err != nil {
		// If it fails, the json is invalid
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	log = log.WithFields(logrus.Fields{
		"type": device.Type,
		"name": device.Name,
	})
	ctx = logger.ToCtx(ctx, log)

	log.Info("Load device")
	// Find the device type in the loader
	loader, err := devices.GetLoader(device.Type)
	if err != nil { // if there was an error while finding the device type
		if err == devices.ErrTypeNotFound {
			// If the device was not found => Invalid request
			utils.Err(ctx, resp, http.StatusBadRequest, "invalid device type")
			log.WithError(err).Error("Invalid device type")
			return nil
		} else {
			// It's another error => 500
			return errors.Wrap(err, "fail to get device loader")
		}
	}

	// We found a loader for this device
	log.Info("Validate")

	// Validate the specialized params
	err = loader.Validate(device.Params)
	if err != nil {
		// If the validation did not pass => Invalid request (the specialization body is invalid)
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		return nil
	}

	log.Info("Save")
	// If everything checked out store this in the database
	err = document.Save(ctx, models.DeviceCollection, &device)
	if err != nil {
		log.WithError(err).Error("Fail to save it")
		return errors.Wrap(err, "fail to save device")
	}

	// Create a new context so the device wont be affected by the request (imeouts, etc.)
	ctx = logger.ToCtx(context.Background(), log)

	log.Info("Add device to the graph")
	// Add the device to the graph
	dev, err := c.graph.Add(ctx, device.ID.Hex())
	if err != nil {
		return errors.Wrap(err, "fail to add device")
	}

	log.Info("Device imported")
	deviceResp := DetailedDeviceResponse{
		DeviceResponse: deviceToDeviceResponse(dev),
		InputPorts:     dev.InputPorts(),
		OutputPorts:    dev.OutputPorts(),
		DisplayOpts:    device.DisplayOpts,
	}
	utils.JSON(ctx, resp, deviceResp)

	return nil
}

func (c DeviceController) Update(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var dev models.Device
	err := document.Find(ctx, models.DeviceCollection, id, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to find device")
	}

	var opts DeviceUpdateParams
	err = json.NewDecoder(req.Body).Decode(&opts)
	if err != nil {
		utils.Err(ctx, resp, http.StatusNotFound, err.Error())
		return nil
	}

	query := bson.M{}
	if opts.DisplayOpts != nil {
		query["display_opts"] = opts.DisplayOpts
	}

	err = document.Update(ctx, models.DeviceCollection, bson.M{"$set": query}, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to update device")
	}
	return nil
}

func (c DeviceController) Destroy(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var dev models.Device
	err := document.Find(ctx, models.DeviceCollection, id, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to find device")
	}

	err = c.graph.Remove(ctx, dev.ID.Hex())
	if err != nil {
		return errors.Wrap(err, "fail to remove device from graph")
	}

	err = document.ReallyDestroy(ctx, models.DeviceCollection, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to remove device from database")
	}
	resp.WriteHeader(http.StatusNoContent)
	return nil
}

func (c DeviceController) APICall(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var dev models.Device
	err := document.Find(ctx, models.DeviceCollection, id, &dev)
	if err != nil {
		return errors.Wrap(err, "fail to find device")
	}

	// Fetch the device in the device graph
	device, err := c.graph.Get(ctx, params["id"])
	if err != nil { // The device graph failed to find the device
		// If the device graph does not exist in the device graph
		if err == graph.ErrNotFound {
			utils.Err(ctx, resp, http.StatusNotFound, "not found")
			return nil
		} else {
			// It's another error => 500
			return errors.Wrap(err, "fail to find device")
		}
	}

	prefix := fmt.Sprintf("/api/devices/%s", params["id"])
	http.StripPrefix(prefix, device.API()).ServeHTTP(resp, req)
	return nil
}

func deviceToDeviceResponse(device devices.Device) DeviceResponse {
	return DeviceResponse{
		ID:    device.ID().Hex(),
		Name:  device.Name(),
		Type:  device.Type(),
		State: device.State(),
	}
}
