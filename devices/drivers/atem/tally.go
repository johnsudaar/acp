package atem

import (
	"context"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/events"
	"github.com/johnsudaar/atem"
	"github.com/sirupsen/logrus"
)

func (a *ATEM) WriteTally(tallies atem.TallyStatuses) {
	log := a.log
	ctx := logger.ToCtx(context.Background(), log)
	for _, tally := range tallies {
		log.WithFields(logrus.Fields{
			"port":    tally.Source.String(),
			"preview": tally.Preview,
			"program": tally.Program,
		})
		log.Info("Send tally")
		a.SendEvent(ctx, tally.Source.String(), events.TallyEventName, events.TallyEvent{
			Program: tally.Program,
			Preview: tally.Preview,
		})
	}
}
