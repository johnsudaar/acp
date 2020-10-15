package tally

import (
	"context"
	"sync"
	"time"

	handlers "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/realtime"
)

type Value string

const (
	Off     Value = "Off"
	Preview Value = "Preview"
	Program Value = "Program"
)

// Check that the TallyDriver is a valid DeviceType
var _ types.DeviceType = &TallyDriver{}

// TODO: It would be nice to have a method on the client side to reconnect tally
type Tallyable interface {
	OutputPorts() []string
	SendTally(ctx context.Context, port string, value Value)
	PublishRealtimeEvent(ctx context.Context, ch realtime.Channel, data interface{})
}

type TallyDriver struct {
	// Needed variables
	device   Tallyable
	values   map[string]Value
	stopping bool

	// Timing stuff
	refreshChan chan bool
	lock        *sync.RWMutex
	stopLock    *sync.Mutex
}

func NewTallyDriver(device Tallyable) *TallyDriver {
	return &TallyDriver{
		device:      device,
		stopping:    false,
		values:      make(map[string]Value),
		refreshChan: make(chan bool, 1),
		lock:        &sync.RWMutex{},
		stopLock:    &sync.Mutex{},
	}
}

func (t *TallyDriver) Start(ctx context.Context) error {
	go t.loop(ctx)
	return nil
}

func (t *TallyDriver) Stop(ctx context.Context) error {
	t.stopLock.Lock()
	t.stopping = true
	t.stopLock.Unlock()
	return nil
}

func (t *TallyDriver) loop(ctx context.Context) {
	log := logger.Get(ctx).WithField("process", "tallydriver")
	for {
		// Are we stopping ?

		t.stopLock.Lock()
		stopping := t.stopping
		t.stopLock.Unlock()
		if stopping {
			log.Info("Stopping")
			return
		}

		log.Info("Refresh tallies")
		t.refreshAllTallies(ctx)

		timer := time.NewTimer(3 * time.Second)

		select {
		case <-timer.C:
		case <-t.refreshChan:
		}
	}
}

func (t *TallyDriver) refreshAllTallies(ctx context.Context) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	for _, port := range t.device.OutputPorts() {
		value, ok := t.values[port]
		if !ok {
			value = Off
		}
		go t.device.SendTally(ctx, port, value)
		go t.device.PublishRealtimeEvent(ctx, realtime.TallyChannel, realtime.TallyEvent{
			Program: value == Program,
			Preview: value == Preview,
		})
	}
}

func (p *TallyDriver) Routes() map[string]handlers.HandlerFunc {
	return nil
}
