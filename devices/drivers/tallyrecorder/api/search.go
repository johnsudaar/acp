package api

import (
	"net/http"

	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/devices/drivers/tallyrecorder/models"
	"github.com/johnsudaar/acp/utils"
	"gopkg.in/mgo.v2/bson"
)

func (a APIHandler) Search(resp http.ResponseWriter, req *http.Request, params map[string]string) error {
	ctx := req.Context()

	shoot := req.URL.Query().Get("shoot")
	if shoot == "" {
		shoot = a.Shoot
	}

	query := bson.M{
		"shoot": shoot,
	}

	log := logger.Default().WithField("shoot", shoot)

	ctx = logger.ToCtx(ctx, log)
	var data []models.TallyEvent
	err := document.Where(ctx, models.TallyEventCollection, query, &data)
	if err != nil {
		return err
	}

	utils.JSON(ctx, resp, data)

	return nil
}
