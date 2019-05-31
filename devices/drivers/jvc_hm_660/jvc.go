package jvc

import (
	"context"
	"sync"
	"time"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/jvc_api/client"
	"github.com/sirupsen/logrus"
)

type JVCHM660 struct {
	*devices.Base
	IP       string
	Port     int
	User     string
	Password string

	log          logrus.FieldLogger
	stoppingLock *sync.Mutex
	stopping     bool

	clientSync *sync.RWMutex
	client     *client.HTTPClient
}

func (j *JVCHM660) InputPorts() []string {
	return []string{}
}

func (j *JVCHM660) OutputPorts() []string {
	return []string{"SDI OUT"}
}

func (j *JVCHM660) Start() error {
	go j.watchDog()
	return nil
}

func (j *JVCHM660) Stop() error {
	j.stoppingLock.Lock()
	defer j.stoppingLock.Unlock()
	j.stopping = true
	return nil
}

func (j *JVCHM660) connect() {
	j.log.Info("Connecting")
	camClient, err := client.New(j.IP, j.User, j.Password)
	if err != nil {
		j.log.WithError(err).Error("fail to connect")
		j.clientSync.Lock()
		j.client = nil
		j.clientSync.Unlock()
		j.SetState(devices.StateNotConnected)
		return
	}

	j.clientSync.Lock()
	j.client = camClient
	err = j.client.SetTallyLampPriority(client.TallyPriorityWeb)
	j.clientSync.Unlock()
	if err != nil {
		j.log.WithError(err).Error("fail to set tally lamp priority")
		return
	}
	j.SetState(devices.StateConnected)
	j.log.Info("Connected")
}

func (j *JVCHM660) Types() []types.Type {
	return []types.Type{types.TallyType}
}
func (j *JVCHM660) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (j *JVCHM660) watchDog() {
	for {
		// Mutex protection
		j.stoppingLock.Lock()
		stopping := j.stopping
		j.stoppingLock.Unlock()
		if stopping {
			j.log.Info("Stopping")
			return
		}

		if j.State() != devices.StateConnected {
			j.connect()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
