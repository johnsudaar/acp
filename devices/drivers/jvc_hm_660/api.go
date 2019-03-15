package jvc

import (
	"io/ioutil"
	"net/http"
	"time"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/utils"
	"github.com/sirupsen/logrus"
)

func (a *JVCHM660) API() http.Handler {
	log := logrus.New()
	log.Out = ioutil.Discard
	router := handlers.NewRouter(log)
	router.HandleFunc("/rec/status", a.RecStatus)
	return router
}

type RecResponse struct {
	Recording     bool   `json:"recording"`
	RecordingTime string `json:"recording_time"`
}

func (a *JVCHM660) RecStatus(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	//if a.State() == devices.StateConnected {
	//	utils.Err(ctx, resp, http.StatusBadRequest, "device not connected")
	//	return nil
	//}

	//status, err := a.client.GetCamStatus()
	//if err != nil {
	//	return errors.Wrap(err, "fail to get cam status")
	//}

	//tcTime, err := status.Camera.TCTime()
	//if err != nil {
	//	return errors.Wrap(err, "fail to get timecode time")
	//}
	//recording := status.Camera.Status == client.CamStatusRecording

	tcTime := 1*time.Hour + 10*time.Minute + 32*time.Second
	recording := true

	recResp := RecResponse{
		RecordingTime: tcTime.String(),
		Recording:     recording,
	}

	utils.JSON(ctx, resp, recResp)

	return nil
}
