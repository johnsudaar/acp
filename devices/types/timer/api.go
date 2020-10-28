package timer

import (
	"net/http"

	"github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
)

func (s *TimerDriver) Routes() map[string]handlers.HandlerFunc {
	return map[string]handlers.HandlerFunc{
		"/timers/sources":            s.Sources,
		"/timers/sources/{tcsource}": s.Value,
	}
}

func (s *TimerDriver) Sources(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	sources := s.device.TimecodeSources()
	utils.JSON(ctx, resp, map[string][]string{
		"sources": sources,
	})
	return nil
}

func (s *TimerDriver) Value(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	tcsource := params["tcsource"]
	if req.Method != http.MethodGet || !utils.HasString(tcsource, s.device.TimecodeSources()) {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"status": "not found"}`))
		return nil
	}

	value, err := s.device.Timecode(tcsource)
	if err != nil {
		return errors.Wrap(err, "fail to get timecode")
	}

	utils.JSON(ctx, resp, value)

	return nil
}
