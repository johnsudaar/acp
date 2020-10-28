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
	"github.com/johnsudaar/acp/realtime"
	"github.com/johnsudaar/acp/scenes"
	"github.com/johnsudaar/acp/timer"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
)

func Start(ctx context.Context, graph graph.Graph, realtime realtime.Realtime, timers *timer.Timers, scenes scenes.Scenes) error {
	log := logger.Get(ctx).WithField("source", "http_server")
	config := config.Get()

	router := handlers.NewRouter(log)
	router.Use(handlers.ErrorMiddleware)

	deviceController := NewDeviceController(graph)
	deviceTypesController := NewDeviceTypesController()
	linkController := NewLinkController(graph)
	timerController := NewTimerController(timers)
	sceneController := NewScenesController(scenes)
	router.HandleFunc("/api/ping", Ping).Methods("GET")
	router.HandleFunc("/api/devices", deviceController.List).Methods("GET")
	router.HandleFunc("/api/devices", deviceController.Create).Methods("POST")
	router.HandleFunc("/api/devices/{id}", deviceController.Show).Methods("GET")
	router.HandleFunc("/api/devices/{id}", deviceController.Destroy).Methods("DELETE")
	router.HandleFunc("/api/devices/{id}", deviceController.Update).Methods("PATCH", "PUT")
	router.HandleFunc("/api/devices/{id}/{_dummy:.*}", deviceController.APICall)
	router.HandleFunc("/api/device_types", deviceTypesController.List).Methods("GET")
	router.HandleFunc("/api/device_types/{id}/params", deviceTypesController.Params).Methods("GET")
	router.HandleFunc("/api/links", linkController.List).Methods("GET")
	router.HandleFunc("/api/links", linkController.Create).Methods("POST")
	router.HandleFunc("/api/links/{id}", linkController.Destroy).Methods("DELETE")
	router.HandleFunc("/api/timers", timerController.List).Methods("GET")
	router.HandleFunc("/api/timers", timerController.Create).Methods("POST")
	router.HandleFunc("/api/timers/{id}", timerController.Update).Methods("PATCH", "PUT")
	router.HandleFunc("/api/timers/{id}", timerController.Destroy).Methods("DELETE")
	router.HandleFunc("/api/timers/{id}/action", timerController.Action).Methods("POST")
	router.HandleFunc("/api/scenes", sceneController.List).Methods("GET")
	router.HandleFunc("/api/scenes", sceneController.Create).Methods("POST")
	router.HandleFunc("/api/scenes/_active", sceneController.Active).Methods("GET")
	router.HandleFunc("/api/scenes/{id}", sceneController.Update).Methods("PATCH", "PUT")
	router.HandleFunc("/api/scenes/{id}", sceneController.Destroy).Methods("DELETE")
	router.HandleFunc("/api/scenes/{id}", sceneController.Show).Methods("GET")
	router.HandleFunc("/api/scenes/{id}/launch", sceneController.Launch).Methods("POST")

	router.Router.PathPrefix("/").HandlerFunc(Front)

	headersOk := muxhandlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"})
	originsOk := muxhandlers.AllowedOrigins([]string{"*"})
	methodsOk := muxhandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	http.HandleFunc("/connection/websocket/", realtime.Websocket)
	http.HandleFunc("/connection/sockjs/", realtime.SockJS)
	http.Handle("/", muxhandlers.CORS(originsOk, headersOk, methodsOk)(router))

	log.WithField("port", config.Server.Port).Info("Starting web server")

	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), nil)
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

func Front(resp http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(config.Get().Server.AssetsPath)).ServeHTTP(resp, req)
}
