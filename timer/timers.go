package timer

import (
	"context"
	"sync"

	"github.com/Scalingo/go-utils/mongo/document"
	"github.com/johnsudaar/acp/graph"
	"github.com/johnsudaar/acp/models"
	"github.com/johnsudaar/acp/realtime"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type Timers struct {
	timers     map[string]*Timer
	realtime   realtime.Realtime
	graph      graph.Graph
	timersLock *sync.RWMutex
}

func Load(ctx context.Context, realtime realtime.Realtime, graph graph.Graph) (*Timers, error) {
	timers := &Timers{
		realtime:   realtime,
		graph:      graph,
		timers:     make(map[string]*Timer),
		timersLock: &sync.RWMutex{},
	}

	var timerModels []models.Timer
	err := document.Where(ctx, models.TimerCollection, bson.M{}, &timerModels)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get timers in database")
	}

	for _, t := range timerModels {
		timers.Add(t)
	}

	return timers, nil
}

func (t *Timers) Add(timer models.Timer) {
	t.timersLock.Lock()
	defer t.timersLock.Unlock()
	t.timers[timer.ID.Hex()] = NewTimer(timer, t.realtime, t.graph)
	t.timers[timer.ID.Hex()].Start()
}

func (t *Timers) Get(id bson.ObjectId) *Timer {
	t.timersLock.RLock()
	defer t.timersLock.RUnlock()

	return t.timers[id.Hex()]
}

func (t *Timers) Remove(id bson.ObjectId) {
	t.timersLock.Lock()
	defer t.timersLock.Unlock()
	timer := t.timers[id.Hex()]
	if timer == nil {
		return
	}

	timer.Stop()
	delete(t.timers, id.Hex())
}
