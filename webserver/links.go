package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/models"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func NewLinkController(g graph.Graph) LinkController {
	if g == nil {
		panic("Graph should not be nil")
	}

	return LinkController{
		graph: g,
	}
}

type LinkController struct {
	graph graph.Graph
}

func (c LinkController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	var links []models.Link
	err := document.Where(ctx, models.LinkCollection, bson.M{}, &links)
	if err != nil {
		return errors.Wrap(err, "fail to get links")
	}

	JSON(ctx, resp, links)
	return nil
}

func (c LinkController) Create(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)

	var link models.Link
	// Decode the link
	err := json.NewDecoder(req.Body).Decode(&link)
	if err != nil {
		// If it fails, the json is invalid
		Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	// Try to connect the devices according to the link
	err = c.graph.Connect(ctx, link.Input, link.Output)
	if err != nil {
		// If it fails, the specified link is impossible
		Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	// Save it in db
	err = document.Save(ctx, models.LinkCollection, &link)
	if err != nil {
		return errors.Wrap(err, "fail to save link")
	}

	// It worked \o/ return the saved link
	JSON(ctx, resp, link)
	return nil
}

func (c LinkController) Destroy(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var link models.Link
	err := document.Find(ctx, models.LinkCollection, id, &link)
	if err != nil {
		return errors.Wrap(err, "fail to find device")
	}

	c.graph.Disconnect(ctx, link.Input, link.Output)

	err = document.ReallyDestroy(ctx, models.LinkCollection, &link)
	if err != nil {
		return errors.Wrap(err, "fail to remove device from database")
	}

	resp.WriteHeader(http.StatusNoContent)
	return nil
}
