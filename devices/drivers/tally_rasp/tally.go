package tally

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Scalingo/go-utils/logger"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/tally"
	"github.com/sirupsen/logrus"
)

// Check if this can work as a tally

var _ tally.Tallyable = &Tally{}

type Tally struct {
	*devices.Base
	IP         string
	log        logrus.FieldLogger
	portsCount int
}

func (t *Tally) Start() error {
	return nil
}

func (t *Tally) Stop() error {
	return nil
}

func (t *Tally) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
}

func (t *Tally) Types() []types.Type {
	return []types.Type{types.TallyType}
}

func (t *Tally) InputPorts() []string {
	return []string{}
}

func (t *Tally) OutputPorts() []string {
	ports := make([]string, t.portsCount)
	for i := 0; i < t.portsCount; i++ {
		ports[i] = strconv.Itoa(i)
	}
	return ports
}

func (t *Tally) SendTally(ctx context.Context, port string, value tally.Value) {
	log := logger.Get(ctx)
	url := fmt.Sprintf("http://%s/tally?tally_id=%s&status=%s", t.IP, port, t.toTallyString(value))
	log.Info(url)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.WithError(err).Error("fail to write tally")
		return
	}
	resp.Body.Close()

}

func (t *Tally) SendTally2(ctx context.Context, port string, value tally.Value) {
	log := logger.Get(ctx)
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(map[string]string{
		"tally_id": port,
		"status":   t.toTallyString(value),
	})

	url := fmt.Sprintf("http://%s/tally", t.IP)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post(url, "application/json", &buff)
	if err != nil {
		log.WithError(err).Error("fail to write tally")
		return
	}
	resp.Body.Close()
}

func (t *Tally) toTallyString(value tally.Value) string {
	switch value {
	case tally.Program:
		return "pgm"
	case tally.Preview:
		return "pvw"
	case tally.Off:
		return "off"
	}

	return "off"
}

func (t *Tally) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}
