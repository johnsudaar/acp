package ptz

import (
	"encoding/json"
	"net/http"

	"github.com/johnsudaar/acp/utils"
)

func (p *PtzDriver) Joystick(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	var payload PTZJoystickParams
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, "invalid json: "+err.Error())
		return nil
	}
	if payload.PanSpeed > 1 {
		payload.PanSpeed = 1
	}
	if payload.PanSpeed < -1 {
		payload.PanSpeed = -1
	}
	if payload.TiltSpeed > 1 {
		payload.TiltSpeed = 1
	}
	if payload.TiltSpeed < -1 {
		payload.TiltSpeed = -1
	}
	if payload.ZoomSpeed > 1 {
		payload.ZoomSpeed = 1
	}
	if payload.ZoomSpeed < -1 {
		payload.ZoomSpeed = -1
	}

	err = p.device.SendPTZJoystick(payload)
	if err != nil {
		utils.Err(ctx, resp, http.StatusInternalServerError, err.Error())
		return nil
	}
	resp.Write([]byte(`{"status": "success"}`))
	return nil
}
