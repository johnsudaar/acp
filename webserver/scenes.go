package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/scenes"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func NewScenesController(scenes scenes.Scenes) ScenesController {
	return ScenesController{
		scenes: scenes,
	}
}

type ScenesController struct {
	scenes scenes.Scenes
}

func (c ScenesController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	var scenes []models.Scene

	err := document.Where(ctx, models.SceneCollection, bson.M{}, &scenes)
	if err != nil {
		return errors.Wrap(err, "fail to get scenes")
	}

	utils.JSON(ctx, resp, scenes)
	return nil
}

func (c ScenesController) Create(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	var scene models.Scene

	err := json.NewDecoder(req.Body).Decode(&scene)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	err = document.Save(ctx, models.SceneCollection, &scene)
	if err != nil {
		log.WithError(err).Error("Fail to save it")
		return errors.Wrap(err, "fail to save scene")
	}

	utils.JSON(ctx, resp, scene)
	return nil
}

func (c ScenesController) Update(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var scene models.Scene
	err := document.Find(ctx, models.SceneCollection, id, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to find timer")
	}

	var opts models.Scene
	err = json.NewDecoder(req.Body).Decode(&opts)
	if err != nil {
		utils.Err(ctx, resp, http.StatusNotFound, err.Error())
		return nil
	}

	scene.Name = opts.Name
	scene.Elements = opts.Elements

	err = document.Save(ctx, models.SceneCollection, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to save scene")
	}
	utils.JSON(ctx, resp, scene)
	return nil
}

func (c ScenesController) Show(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var scene models.Scene
	err := document.Find(ctx, models.SceneCollection, id, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to find scene")
	}

	utils.JSON(ctx, resp, scene)
	return nil
}

func (c ScenesController) Destroy(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var scene models.Scene
	err := document.Find(ctx, models.SceneCollection, id, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to find scene")
	}

	err = document.ReallyDestroy(ctx, models.SceneCollection, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to remove scene from database")
	}

	resp.WriteHeader(http.StatusNoContent)
	return nil
}

func (c ScenesController) Launch(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var scene models.Scene
	err := document.Find(ctx, models.SceneCollection, id, &scene)
	if err != nil {
		return errors.Wrap(err, "fail to find scene")
	}

	c.scenes.SetActive(id.Hex())
	resp.WriteHeader(http.StatusNoContent)
	return nil
}

func (c ScenesController) Active(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	id := c.scenes.Active()

	utils.JSON(ctx, resp, map[string]string{
		"scene_id": id,
	})
	return nil
}
