package timer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/johnsudaar/acp/devices/types"
)

type TimerType string

const (
	TimerTypeTime     TimerType = "time"
	TimerTypeDuration TimerType = "duration"
	TimerTypeNA       TimerType = "n/a"
)

type TimerValue struct {
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
	Type     TimerType     `json:"type"`
}

func NewDurationValue(d time.Duration) TimerValue {
	return TimerValue{
		Duration: d,
		Type:     TimerTypeDuration,
	}
}
func NewNAValue() TimerValue {
	return TimerValue{
		Type: TimerTypeNA,
	}
}

type Timeable interface {
	TimecodeSources() []string
	Timecode(source string) (TimerValue, error)
}

var _ types.DeviceType = &TimerDriver{}

type TimerDriver struct {
	device Timeable
}

func NewTimerDriver(device Timeable) *TimerDriver {
	return &TimerDriver{
		device: device,
	}
}

func (s *TimerDriver) Start(ctx context.Context) error {
	return nil
}

func (s *TimerDriver) Stop(ctx context.Context) error {
	return nil
}

func (p *TimerDriver) EventSubscriptions() []string {
	return []string{}
}

func (p *TimerDriver) RealtimeEventSubscriptions() []string {
	return []string{}
}

func (p *TimerDriver) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (p *TimerDriver) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}
