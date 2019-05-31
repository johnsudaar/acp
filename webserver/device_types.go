package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/utils"
)

func NewDeviceTypesController() DeviceTypesController {
	return DeviceTypesController{}
}

type DeviceTypesController struct {
}

func (DeviceTypesController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	types := devices.AvailableTypes()

	utils.JSON(ctx, resp, types)
	return nil
}

func (DeviceTypesController) Params(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	id := params["id"]

	if id == "" {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}

	loader, err := devices.GetLoader(id)
	if err != nil {
		if err == devices.ErrTypeNotFound {
			utils.Err(ctx, resp, http.StatusNotFound, "not found")
			return nil
		}

		return err
	}

	deviceParams := loader.Params()

	err = json.NewEncoder(resp).Encode(deviceParams)
	if err != nil {
		log.WithError(err).Error("Fail to encode body")
	}

	return nil
}
