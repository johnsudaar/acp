package x32

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices/drivers/x32/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type APIHandler struct{}

type RoomParams struct {
	Name    *string `json:"name"`
	Channel *int    `json:"channel"`
	Mix     *int    `json:"mix"`
}

func (a APIHandler) ListRooms(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)

	var rooms []models.Room
	err := document.Where(ctx, models.X32IntercomRooms, bson.M{}, &rooms)
	if err != nil {
		return errors.Wrap(err, "fail to find room")
	}

	err = json.NewEncoder(resp).Encode(map[string][]models.Room{
		"rooms": rooms,
	})
	if err != nil {
		log.WithError(err).Error("fail to encode rooms")
	}
	return nil
}

func (a APIHandler) AddRoom(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)

	var body RoomParams

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Infof("Fail to decode params %v", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{"error": err.Error()})
		return nil
	}

	room := models.Room{
		Name:    *body.Name,
		Channel: *body.Channel,
		Mix:     *body.Mix,
	}

	err = document.Create(ctx, models.X32IntercomRooms, &room)
	if err != nil {
		return errors.Wrap(err, "fail to save room")
	}

	err = json.NewEncoder(resp).Encode(map[string]models.Room{
		"room": room,
	})
	if err != nil {
		log.WithError(err).Error("fai to encode room")
	}

	return nil
}

func (a APIHandler) UpdateRoom(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()
	log := logger.Get(ctx)

	var body RoomParams

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Infof("Fail to decode params %v", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{"error": err.Error()})
		return nil
	}

	roomID := params["room_id"]
	if !bson.IsObjectIdHex(roomID) {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{"error": "not found"})
		return nil
	}

	var room models.Room
	err = document.Find(ctx, models.X32IntercomRooms, bson.ObjectIdHex(roomID), &room)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteHeader(http.StatusNotFound)
			json.NewEncoder(resp).Encode(map[string]string{"error": "not found"})
			return nil
		}
		return errors.Wrap(err, "fail to find room")
	}

	if body.Name != nil {
		room.Name = *body.Name
	}
	if body.Channel != nil {
		room.Channel = *body.Channel
	}

	if body.Mix != nil {
		room.Mix = *body.Mix
	}

	err = document.Save(ctx, models.X32IntercomRooms, &room)
	if err != nil {
		return errors.Wrap(err, "fail to save room")
	}

	err = json.NewEncoder(resp).Encode(map[string]models.Room{
		"room": room,
	})
	if err != nil {
		log.WithError(err).Error("fai to encode room")
	}

	return nil
}

func (a APIHandler) DeleteRoom(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	roomID := params["room_id"]
	if !bson.IsObjectIdHex(roomID) {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{"error": "not found"})
		return nil
	}

	var room models.Room
	err := document.Find(ctx, models.X32IntercomRooms, bson.ObjectIdHex(roomID), &room)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteHeader(http.StatusNotFound)
			json.NewEncoder(resp).Encode(map[string]string{"error": "not found"})
			return nil
		}
		return errors.Wrap(err, "fail to find room")
	}

	err = document.Destroy(ctx, models.X32IntercomRooms, &room)
	if err != nil {
		return errors.Wrap(err, "fail to destroy room")
	}
	return nil
}

func (a APIHandler) RoomOn(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	return nil
}

func (a APIHandler) RoomOff(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	return nil
}

func (s *X32) API() http.Handler {
	apiHandler := APIHandler{}
	log := logrus.New()
	log.Out = ioutil.Discard
	router := handlers.NewRouter(log)
	router.HandleFunc("/rooms", apiHandler.ListRooms).Methods("GET")
	router.HandleFunc("/rooms", apiHandler.AddRoom).Methods("POST")
	router.HandleFunc("/rooms/{room_id}", apiHandler.UpdateRoom).Methods("PATCH", "PUT")
	router.HandleFunc("/rooms/{room_id}", apiHandler.DeleteRoom).Methods("DELETE")
	router.HandleFunc("/rooms/{room_id}/input/{input_id}/on", apiHandler.RoomOn).Methods("POST")
	router.HandleFunc("/rooms/{room_id}/input/{input_id}/off", apiHandler.RoomOff).Methods("POST")
	return router
}
