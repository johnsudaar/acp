package smartview

import (
	"context"
	"fmt"
	"net"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/tally"
	"github.com/sirupsen/logrus"
)

type SmartView struct {
	*devices.Base
	IP      string
	Outputs []string
	log     logrus.FieldLogger
	client  net.Conn
}

func (t *SmartView) Start() error {
	go t.connect()
	return nil
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
	return nil
}

func (t *SmartView) InputPorts() []string {
	return []string{}
}

func (t *SmartView) OutputPorts() []string {
	return t.Outputs
}

func (t *SmartView) Types() []types.Type {
	return []types.Type{types.TallyType}
}
func (t *SmartView) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (t *SmartView) SendTally(ctx context.Context, port string, value tally.Value) {
	log := logger.Get(ctx)
	if t.client == nil {
		log.Error("Client not initialized")
		return
	}

	payload := fmt.Sprintf("%s:\nBorder: %s\n\n", port, t.toTallyString(value))

	_, err := t.client.Write([]byte(payload))
	if err != nil {
		log.WithError(err).Error("fail to send command to smartview")
		return
	}
}
func (t *SmartView) toTallyString(value tally.Value) string {
	switch value {
	case tally.Program:
		return "RED"
	case tally.Preview:
		return "GREEN"
	}
	return "NONE"
}
