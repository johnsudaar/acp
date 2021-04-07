package atem

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/atem"
	"github.com/sirupsen/logrus"
)

type atemIOConfig struct {
	Inputs int
}

var ioConfigs map[string]atemIOConfig = map[string]atemIOConfig{
	AtemType1ME: {
		Inputs: 10,
	},
	AtemTypeConstellationHD: {
		Inputs: 40,
	},
}

type ATEM struct {
	*devices.Base
	IP   string
	Port string

	log          logrus.FieldLogger
	stoppingLock *sync.Mutex
	stopping     bool

	client *atem.AtemClient

	atemType string
}

func (a *ATEM) Start() error {
	go a.watchDog()

	return nil
}

func (a *ATEM) connect() {
	if a.client != nil {
		a.client.Close()
	}
	ctx := logger.ToCtx(context.Background(), a.log)

	a.log.Info("Trying to connect")
	client, err := atem.New(ctx,
		net.JoinHostPort(a.IP, a.Port),
		atem.WithTallyWriter(a),
	)
	if err != nil {
		a.client = nil
		a.SetState(devices.StateNotConnected)
		a.log.WithError(err).Error("Fail to connect")
		return
	}

	a.log.Info("Connected")
	a.client = client
	a.SetState(devices.StateConnected)
}

func (a *ATEM) Stop() error {
	a.stoppingLock.Lock()
	defer a.stoppingLock.Unlock()
	a.stopping = true
	return nil
}

func (a *ATEM) watchDog() {
	log := a.log
	startTime := time.Now()
	for {
		// Mutex protection
		a.stoppingLock.Lock()
		stopping := a.stopping
		a.stoppingLock.Unlock()

		if stopping {
			log.Info("Stopping")
			if a.client != nil {
				a.client.Close()
				a.client = nil
			}
			log.Info("Stopped")
			return
		}

		if a.State() != devices.StateConnected || time.Now().Sub(startTime) > 3*time.Minute {
			log.Info("Reconnecting")
			a.connect()
			startTime = time.Now()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (a *ATEM) Types() []types.Type {
	return []types.Type{
		types.SwitcherType,
	}
}
