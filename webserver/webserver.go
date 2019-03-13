package webserver

import (
	"context"
	"fmt"
	"net/http"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/johnsudaar/acp/config"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
)

func Start(ctx context.Context, graph graph.Graph) error {
	log := logger.Get(ctx).WithField("source", "http_server")
	config := config.Get()

	router := handlers.NewRouter(log)

	deviceController := NewDeviceController(graph)
	linkController := NewLinkController(graph)
	router.HandleFunc("/api/ping", Ping).Methods("GET")
	router.HandleFunc("/api/devices", deviceController.List).Methods("GET")
	router.HandleFunc("/api/devices", deviceController.Create).Methods("POST")
	router.HandleFunc("/api/devices/{id}", deviceController.Show).Methods("GET")
	router.HandleFunc("/api/devices/{id}", deviceController.Destroy).Methods("DELETE")
	router.HandleFunc("/api/devices/{id}", deviceController.Update).Methods("PATCH", "PUT")
	router.HandleFunc("/api/devices/{id}/{_dummy:.*}", deviceController.APICall)
	router.HandleFunc("/api/device_types", deviceController.ListTypes).Methods("GET")
	router.HandleFunc("/api/links", linkController.List).Methods("GET")
	router.HandleFunc("/api/links", linkController.Create).Methods("POST")
	router.HandleFunc("/api/links/{id}", linkController.Destroy).Methods("DELETE")

	headersOk := muxhandlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"})
	originsOk := muxhandlers.AllowedOrigins([]string{"*"})
	methodsOk := muxhandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELEtE"})

	log.WithField("port", config.Port).Info("Starting web server")
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Port), muxhandlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		return errors.Wrap(err, "HTTP server failed")
	}
	return nil
}

func Ping(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	utils.JSON(req.Context(), resp, map[string]string{
		"response": "pong",
	})

	return nil
}
