package switcher

import (
	"context"
	"encoding/json"
	"net/http"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/utils"
)

var _ types.DeviceType = &SwitcherDriver{}

type SwitcherDriver struct {
	device Switchable
}

type Switchable interface {
	Switch(output, input string) error
}

func NewSwitcherDriver(device Switchable) *SwitcherDriver {
	return &SwitcherDriver{
		device: device,
	}
}

func (s *SwitcherDriver) Start(ctx context.Context) error {
	return nil
}

func (s *SwitcherDriver) Stop(ctx context.Context) error {
	return nil
}

func (p *SwitcherDriver) EventSubscriptions() []string {
	return []string{}
}

func (p *SwitcherDriver) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (s *SwitcherDriver) Routes() map[string]handlers.HandlerFunc {
	return map[string]handlers.HandlerFunc{
		"/switcher/switch": s.Switch,
	}
}

type SwitchParams struct {
	Output string `json:"output"`
	Input  string `json:"input"`
}

func (s *SwitcherDriver) Switch(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	var payload SwitchParams
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, "invalid json: "+err.Error())
		return nil
	}

	err = s.device.Switch(payload.Output, payload.Input)
	if err != nil {
		log.WithError(err).Error("fail to send switch")
		utils.Err(ctx, resp, http.StatusInternalServerError, err.Error())
		return nil
	}
	resp.Write([]byte(`{"status": "success"}`))
	return nil
}
