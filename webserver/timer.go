package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/timer"
	"github.com/johnsudaar/acp/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type TimerUpdateOpts struct {
	Name           *string           `json:"name"`
	Type           *models.TimerType `json:"type"`
	Duration       *models.Duration  `json:"duration"`
	ExternalDevice *bson.ObjectId    `json:"external_device"`
	ExternalSource *string           `json:"external_source"`
}

func NewTimerController(timers *timer.Timers) TimerController {
	return TimerController{
		timers: timers,
	}
}

type TimerController struct {
	timers *timer.Timers
}

func (c TimerController) List(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	var timers []models.Timer

	err := document.Where(ctx, models.TimerCollection, bson.M{}, &timers)
	if err != nil {
		return errors.Wrap(err, "fail to get timers")
	}
	utils.JSON(ctx, resp, timers)
	return nil
}

func (c TimerController) Create(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)
	var timer models.Timer

	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		// If it fails, the json is invalid
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		log.WithError(err).Error("Invalid body")
		return nil
	}

	log = log.WithFields(logrus.Fields{
		"type": timer.Type,
	})

	log.Info("Save timer")

	err = document.Save(ctx, models.TimerCollection, &timer)
	if err != nil {
		log.WithError(err).Error("Fail to save it")
		return errors.Wrap(err, "fail to save timer")
	}

	log.Info("Load timer")

	c.timers.Add(timer)

	utils.JSON(ctx, resp, timer)

	return nil
}

func (c TimerController) Update(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var timer models.Timer
	err := document.Find(ctx, models.TimerCollection, id, &timer)
	if err != nil {
		return errors.Wrap(err, "fail to find timer")
	}

	var opts TimerUpdateOpts
	err = json.NewDecoder(req.Body).Decode(&opts)
	if err != nil {
		utils.Err(ctx, resp, http.StatusNotFound, err.Error())
		return nil
	}

	query := bson.M{}
	if opts.Type != nil {
		query["type"] = *opts.Type
		timer.Type = *opts.Type
	}
	if opts.Duration != nil {
		query["duration"] = *opts.Duration
		timer.Duration = *opts.Duration
	}
	if opts.Name != nil {
		query["name"] = *opts.Name
		timer.Name = *opts.Name
	}

	if opts.ExternalDevice != nil && len(*opts.ExternalDevice) != 0 {
		query["external_device"] = *opts.ExternalDevice
		timer.ExternalDevice = *opts.ExternalDevice
	}
	if opts.ExternalSource != nil {
		query["external_source"] = *opts.ExternalSource
		timer.ExternalSource = *opts.ExternalSource
	}

	err = document.Update(ctx, models.TimerCollection, bson.M{"$set": query}, &timer)
	if err != nil {
		return errors.Wrap(err, "fail to update timer")
	}

	timerDriver := c.timers.Get(id)
	if timerDriver == nil {
		return errors.New("timer is not loaded")
	}

	timerDriver.UpdateTimer(timer)

	return nil
}

func (c TimerController) Destroy(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	var timer models.Timer
	err := document.Find(ctx, models.TimerCollection, id, &timer)
	if err != nil {
		return errors.Wrap(err, "fail to find timer")
	}

	c.timers.Remove(id)
	err = document.ReallyDestroy(ctx, models.TimerCollection, &timer)
	if err != nil {
		return errors.Wrap(err, "fail to remove timer from database")
	}

	resp.WriteHeader(http.StatusNoContent)
	return nil
}

func (c TimerController) Action(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	if !bson.IsObjectIdHex(params["id"]) {
		utils.Err(ctx, resp, http.StatusNotFound, "not found")
		return nil
	}
	id := bson.ObjectIdHex(params["id"])
	timer := c.timers.Get(id)
	if timer == nil {
		return errors.New("fail to find timer")
	}

	var action models.TimerAction
	err := json.NewDecoder(req.Body).Decode(&action)
	if err != nil {
		utils.Err(ctx, resp, http.StatusBadRequest, err.Error())
		return nil
	}

	err = timer.Do(action)
	if err != nil {
		return errors.Wrap(err, "fail to perform action")
	}
	return nil
}
