package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func NewPositionGroupsController() PositionGroupsController {
	return PositionGroupsController{}
}

type PositionGroupsController struct {
}

func (c PositionGroupsController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	var groups []models.PositionGroup
	err := document.Where(ctx, models.PositionGroupCollection, bson.M{}, &groups)
	if err != nil {
		return errors.Wrap(err, "fail to list groups")
	}
	utils.JSON(ctx, resp, groups)
	return nil
}

func (c PositionGroupsController) Create(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)

	var group models.PositionGroup

	err := json.NewDecoder(req.Body).Decode(&group)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	err = document.Save(ctx, models.PositionGroupCollection, &group)
	if err != nil {
		log.WithError(err).Error("Fail to save position group")
		return errors.Wrap(err, "fail to save position group")
	}

	utils.JSON(ctx, resp, group)
	return nil
}

func (c PositionGroupsController) Destroy(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}

	id := bson.ObjectIdHex(params["id"])

	var group models.PositionGroup
	err := document.Find(ctx, models.PositionGroupCollection, id, &group)
	if err != nil {
		return errors.Wrap(err, "fail to find position group")
	}

	var positions []models.PtzPosition
	err = document.Where(ctx, models.PtzPositionCollection, bson.M{
		"position_group_id": id,
	}, &positions)

	if err != nil {
		return errors.Wrap(err, "fail to check if there are some position linked to this group")
	}

	if len(positions) != 0 {
		utils.Err(ctx, resp, http.StatusBadRequest, "this group is still in use")
		return nil
	}

	err = document.ReallyDestroy(ctx, models.PositionGroupCollection, &group)
	if err != nil {
		return errors.Wrap(err, "fail to remove position group from database")
	}
	resp.WriteHeader(http.StatusNoContent)
	return nil
}
