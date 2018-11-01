package main

import (
	"net/http"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/webserver"
)

func main() {
	log := logger.Default()
	graph, err := graph.Load("./resources/network.json")
	if err != nil {
		panic(err)
	}

	router := handlers.NewRouter(log)
	router.HandleFunc("/api/graph", webserver.GraphController{Graph: graph}.Show).Methods("GET")
	router.HandleFunc("/api/devices/{id}", webserver.DeviceController{Graph: graph}.Show).Methods("GET")

	http.ListenAndServe(":8080", router)

}
