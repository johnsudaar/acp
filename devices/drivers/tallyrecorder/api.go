package tallyrecorder

import (
	"io/ioutil"
	"net/http"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/johnsudaar/acp/devices/drivers/tallyrecorder/api"
	"github.com/sirupsen/logrus"
)

func (a *Recorder) API() http.Handler {
	apiHandler := api.APIHandler{
		Shoot: a.Shoot,
	}
	log := logrus.New()
	log.Out = ioutil.Discard
	router := handlers.NewRouter(log)
	router.HandleFunc("/search", apiHandler.Search)
	return router
}
