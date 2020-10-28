package hyperdeck

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/Scalingo/go-utils/logger"
	"github.com/a-contre-plongee/hyperdeck"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/timer"
	"github.com/sirupsen/logrus"
)

var _ timer.Timeable = &Hyperdeck{}

type Hyperdeck struct {
	*devices.Base
	IP           string
	RepeaterIP   string
	log          logrus.FieldLogger
	client       *hyperdeck.Client
	clientLock   *sync.RWMutex
	stopping     bool
	stoppingLock *sync.Mutex

	cacheLock       *sync.Mutex
	clipCache       map[int]hyperdeck.Clip
	currentClip     int
	timecode        hyperdeck.Timecode
	lastRefreshedAt time.Time
	playing         bool
}

func (h *Hyperdeck) Start() error {
	go h.watchDog()
	go h.timers()
	return nil
}

func (h *Hyperdeck) Stop() error {
	h.stoppingLock.Lock()
	defer h.stoppingLock.Unlock()
	h.stopping = true
	return nil
}

func (h *Hyperdeck) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (h *Hyperdeck) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}

func (h *Hyperdeck) Types() []types.Type {
	return []types.Type{types.TimerType}
}

func (h *Hyperdeck) InputPorts() []string {
	return []string{
		"Input",
	}
}

func (h *Hyperdeck) OutputPorts() []string {
	return []string{
		"Output",
	}
}

func (h *Hyperdeck) watchDog() {
	log := h.log
	for {
		// Mutex protection
		h.stoppingLock.Lock()
		stopping := h.stopping
		h.stoppingLock.Unlock()

		if stopping {
			log.Info("Stopping")
			if h.client != nil {
				h.clientLock.Lock()
				h.client.Stop()
				h.client = nil
				h.clientLock.Unlock()
			}
			log.Info("Stopped")
			return
		}

		client := h.Client()

		if client == nil || !client.Running() {
			h.SetState(devices.StateNotConnected)
		} else {
			h.SetState(devices.StateConnected)
		}

		if h.State() != devices.StateConnected {
			log.Info("Reconnecting")
			h.connect()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (h *Hyperdeck) connect() {
	log := h.log
	if h.Client() != nil {
		h.Client().Stop()
	}

	log.Info("New")

	opts := make([]hyperdeck.ClientOpts, 0)
	if h.RepeaterIP != "" {
		opts = append(opts, hyperdeck.WithRepeater(h.RepeaterIP))
	}

	client := hyperdeck.New(h.IP, opts...)

	log.Info("Start")
	ctx := logger.ToCtx(context.Background(), h.log)
	err := client.Start(ctx)
	log.Info("Started")
	if err != nil {
		h.client = nil
		h.log.WithError(err).Error("fail to start hyperdeck client")
		return
	}
	h.clientLock.Lock()
	h.client = client
	h.clientLock.Unlock()
	log.Info("Success")
}

func (h *Hyperdeck) Client() *hyperdeck.Client {
	h.clientLock.RLock()
	defer h.clientLock.RUnlock()
	return h.client
}
