package tally

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Tally struct {
	*devices.Base
	IP               string
	tallyValue       string
	tallySync        *sync.RWMutex
	tallyRefreshChan chan bool
	log              logrus.FieldLogger
	stopping         bool
	stoppingLock     *sync.Mutex
}

func (t *Tally) Start() error {
	go t.tallyLoop()
	return nil
}

func (t *Tally) Stop() error {
	t.stoppingLock.Lock()
	t.stopping = true
	t.stoppingLock.Unlock()
	return nil
}

func (t *Tally) InputPorts() []string {
	return []string{}
}

func (t *Tally) OutputPorts() []string {
	return []string{"Tally"}
}

func (t *Tally) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	if name == events.TallyEventName {
		params, ok := data.(events.TallyEvent)
		if !ok {
			t.log.Error("Invalid data type for tally event")
		}

		tallyValue := "off"
		if params.Program {
			tallyValue = "pgm"
		} else if params.Preview {
			tallyValue = "pvw"
		}
		t.log.WithField("tally", tallyValue).Info("SetTally")
		t.tallySync.Lock()
		t.tallyValue = tallyValue
		t.tallySync.Unlock()
		t.tallyRefreshChan <- true
	}
}

func (t *Tally) sendTally(value string) error {
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(map[string]string{
		"status": value,
	})
	if err != nil {
		return errors.Wrap(err, "fail to encode tally command")
	}
	for i := 0; i < 3; i++ {
		resp, err := http.Get(fmt.Sprintf("http://%s/tally?tally_id=%v&status=%s", t.IP, i, value))
		if err != nil {
			//log.WithError(err).Error("fail to send tally")
		}
		defer resp.Body.Close()
	}
	return nil
}

func (t *Tally) tallyLoop() {
	log := t.log.WithField("source", "tallyLoop")
	for {
		// Are we stopping?

		t.stoppingLock.Lock()
		stopping := t.stopping
		t.stoppingLock.Unlock()
		if stopping {
			return
		}

		// Get current tally value
		t.tallySync.RLock()
		tallyValue := t.tallyValue
		t.tallySync.RUnlock()

		if tallyValue != "" {
			log.WithField("status", tallyValue).Info("Refresh tally")
			err := t.sendTally(tallyValue)
			if err != nil {
				log.WithError(err).Error("fail to refresh tally")
			}
		}

		timer := time.NewTimer(3 * time.Second)
		select {
		case <-timer.C:
		case <-t.tallyRefreshChan:
		}
	}
}
