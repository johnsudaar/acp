package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/graph"
)

type DeviceController struct {
	Graph graph.Graph
}

func (g DeviceController) Show(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	log := logger.Get(req.Context())

	resp.Header().Set("Content-Type", "application/json")
	id := params["id"]
	device, ok := g.Graph.Devices[id]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"error": "not found"}`))
		return nil
	}

	err := json.NewEncoder(resp).Encode(device.Device)
	if err != nil {
		log.WithError(err).Error("fail to encode graph")
	}
	return nil
}
