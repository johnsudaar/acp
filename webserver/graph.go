package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/graph"
)

type GraphController struct {
	Graph graph.Graph
}

func (g GraphController) Show(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	log := logger.Get(req.Context())

	resp.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(resp).Encode(g.Graph)
	if err != nil {
		log.WithError(err).Error("fail to encode graph")
	}
	return nil
}
