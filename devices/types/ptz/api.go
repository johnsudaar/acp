package ptz

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/utils"
	"gopkg.in/mgo.v2/bson"
)

func (p *PtzDriver) Joystick(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.Err(ctx, resp, http.StatusInternalServerError, "invalid json: "+err.Error())
		log.WithError(err).Error("Fail to send PTZ")
		return nil
	}

	err = p.JoystickAction(body)
	if err != nil {
		utils.Err(ctx, resp, http.StatusInternalServerError, "invalid json: "+err.Error())
		return nil
	}

	resp.Write([]byte(`{"status": "success"}`))
	return nil
}

func (p *PtzDriver) Position(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	var payload PTZPositionParams
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, "invalid json: "+err.Error())
		return nil
	}

	if payload.PositionID != "" {
		var pos models.PtzPosition
		if !bson.IsObjectIdHex(payload.PositionID) {
			utils.Err(ctx, resp, http.StatusNotFound, "invalid id")
			return nil
		}
		err := document.Find(ctx, models.PtzPositionCollection, bson.ObjectIdHex(payload.PositionID), &pos)
		if err != nil {
			utils.Err(ctx, resp, http.StatusNotFound, "not found: "+err.Error())
			return nil
		}
		payload.Pan = pos.Pan
		payload.Tilt = pos.Tilt
		payload.Zoom = pos.Zoom
		payload.Focus = pos.Focus
	}

	err = p.device.SetPosition(payload)
	if err != nil {
		log.WithError(err).Error("fail to send ptz")
		utils.Err(ctx, resp, http.StatusInternalServerError, err.Error())
		return nil
	}
	resp.Write([]byte(`{"status": "success"}`))

	return nil
}

func (p *PtzDriver) Store(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if req.Method == http.MethodPost {
		var payload models.PtzPosition
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.Err(ctx, resp, http.StatusBadRequest, "invalid json: "+err.Error())
			return nil
		}

		if payload.PositionGroupID != nil {
			var pg models.PositionGroup
			document.Find(ctx, models.PositionGroupCollection, *payload.PositionGroupID, &pg)
			if err != nil {
				utils.Err(ctx, resp, http.StatusBadRequest, "invalid position_group_id: "+err.Error())
				return nil
			}
		}

		payload.DeviceID = p.device.ID()

		err = document.Create(ctx, models.PtzPositionCollection, &payload)
		if err != nil {
			utils.Err(ctx, resp, http.StatusInternalServerError, "fail to save position: "+err.Error())
			return nil
		}

		json.NewEncoder(resp).Encode(payload)
		return nil
	} else if req.Method == http.MethodGet {
		var positions []models.PtzPosition

		err := document.Where(ctx, models.PtzPositionCollection, bson.M{
			"device_id": p.device.ID(),
		}, &positions)
		if err != nil {
			utils.Err(ctx, resp, http.StatusInternalServerError, "fail to find positions: "+err.Error())
		}

		json.NewEncoder(resp).Encode(positions)
		return nil
	}
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"status": "not found"}`))
	return nil
}

func (p *PtzDriver) SinglePosition(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["position"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}

	if req.Method == http.MethodDelete {
		var pos models.PtzPosition
		err := document.Find(ctx, models.PtzPositionCollection, bson.ObjectIdHex(params["position"]), &pos)
		if err != nil {
			utils.Err(ctx, resp, http.StatusNotFound, "not found: "+err.Error())
			return nil
		}

		err = document.ReallyDestroy(ctx, models.PtzPositionCollection, &pos)
		if err != nil {
			utils.Err(ctx, resp, http.StatusInternalServerError, "fail to delete position: "+err.Error())
		}
		resp.WriteHeader(http.StatusNoContent)
		return nil
	} else if req.Method == http.MethodPut {
		var payload, pos models.PtzPosition

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.Err(ctx, resp, http.StatusBadRequest, "invalid json: "+err.Error())
			return nil
		}

		err = document.Find(ctx, models.PtzPositionCollection, bson.ObjectIdHex(params["position"]), &pos)
		if err != nil {
			utils.Err(ctx, resp, http.StatusNotFound, "not found: "+err.Error())
			return nil
		}

		if payload.PositionGroupID != nil {
			var pg models.PositionGroup
			document.Find(ctx, models.PositionGroupCollection, *payload.PositionGroupID, &pg)
			if err != nil {
				utils.Err(ctx, resp, http.StatusBadRequest, "invalid position_group_id: "+err.Error())
				return nil
			}
		}

		pos.Name = payload.Name
		pos.Pan = payload.Pan
		pos.Tilt = payload.Tilt
		pos.Zoom = payload.Zoom
		pos.Focus = payload.Focus
		pos.PositionGroupID = payload.PositionGroupID

		err = document.Update(ctx, models.PtzPositionCollection, bson.M{
			"$set": bson.M{
				"name":              pos.Name,
				"pan":               pos.Pan,
				"tilt":              pos.Tilt,
				"focus":             pos.Focus,
				"zoom":              pos.Zoom,
				"position_group_id": pos.PositionGroupID,
			},
		}, &pos)
		if err != nil {
			utils.Err(ctx, resp, http.StatusInternalServerError, "fail to store position: "+err.Error())
			return nil
		}
		json.NewEncoder(resp).Encode(&pos)
		return nil
	}
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"status": "not found"}`))
	return nil

}
