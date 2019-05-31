package smartview

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type SmartView struct {
	*devices.Base
	IP               string
	tallyValues      map[string]string
	tallySync        *sync.RWMutex
	tallyRefreshChan chan bool
	log              logrus.FieldLogger
	stopping         bool
	stoppingLock     *sync.Mutex
	client           net.Conn
}

func (t *SmartView) Start() error {
	go t.tallyLoop()
	//go t.watchDog()
	go t.connect()
	return nil
}

func (t *SmartView) watchDog() {
	log := t.log
	for {
		// Mutex protection
		t.stoppingLock.Lock()
		stopping := t.stopping
		t.stoppingLock.Unlock()

		if stopping {
			log.Info("Stopping")
			if t.client != nil {
				t.client.Close()
				t.client = nil
			}
			log.Info("Stopped")
			return
		}

		if t.State() != devices.StateConnected {
			log.Info("Reconnecting")
			t.connect()
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (t *SmartView) connect() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:9992", t.IP))
	if err != nil {
		t.log.Error("Fail to connect to SmartviewDuo")
		return
	}
	if t.client != nil {
		t.client.Close()
	}
	t.client = conn
}

func (t *SmartView) listen(conn net.Conn) {
	buffer := make([]byte, 4096)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			t.SetState(devices.StateNotConnected)
			return
		}
		t.SetState(devices.StateConnected)
	}
}

func (t *SmartView) Stop() error {
	t.stoppingLock.Lock()
	t.stopping = true
	t.stoppingLock.Unlock()
	return nil
}

func (t *SmartView) InputPorts() []string {
	return []string{}
}

func (t *SmartView) OutputPorts() []string {
	var res []string
	t.tallySync.RLock()
	for k, _ := range t.tallyValues {
		res = append(res, k)
	}
	t.tallySync.RUnlock()
	return res
}

func (t *SmartView) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	if name == events.TallyEventName {
		params, ok := data.(events.TallyEvent)
		if !ok {
			t.log.Error("Invalid data type for tally event")
		}

		tallyValue := "NONE"
		if params.Program {
			tallyValue = "RED"
		} else if params.Preview {
			tallyValue = "GREEN"
		}
		t.log.WithField("tally", tallyValue).Info("SetTally")
		t.tallySync.Lock()
		if _, ok := t.tallyValues[toPort]; !ok {
			t.log.WithField("port", toPort).Error("Port not found")
		} else {
			t.tallyValues[toPort] = tallyValue
		}
		t.tallySync.Unlock()
		t.tallyRefreshChan <- true
	}
}

func (t *SmartView) sendTally(port, value string) error {
	if t.client == nil {
		return errors.New("Not connected")
	}

	payload := fmt.Sprintf("%s\nBorder: %s\n\n", port, value)

	_, err := t.client.Write([]byte(payload))
	if err != nil {
		return errors.Wrap(err, "fail to write to tally")
	}

	return nil
}

func (t *SmartView) tallyLoop() {
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
		for port, value := range t.tallyValues {
			log.WithFields(logrus.Fields{
				"port":  port,
				"value": value,
			}).Info("Refresh tally")

			go func() {
				err := t.sendTally(port, value)
				if err != nil {
					log.WithError(err).Error("fail to refresh tallyA")
				}
			}()

		}
		t.tallySync.RUnlock()

		timer := time.NewTimer(3 * time.Second)
		select {
		case <-timer.C:
		case <-t.tallyRefreshChan:
		}
	}
}
