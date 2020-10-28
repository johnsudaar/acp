package timer

import (
	"context"

	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types/timer"
)

func (t *Timer) ExternalDuration() string {
	ctx := context.Background()
	t.timerLock.RLock()
	deviceId := t.timer.ExternalDevice.Hex()
	source := t.timer.ExternalSource
	t.timerLock.RUnlock()

	device, err := t.graph.Get(ctx, deviceId)
	if err != nil {
		return "-EE:EE:EE"
	}

	wrap, ok := device.(*devices.DeviceWrapper)
	if ok {
		device = wrap.Implementation
	}

	timeable, ok := device.(timer.Timeable)
	if !ok {
		return "-EE:00:EE"
	}

	res, err := timeable.Timecode(source)
	if err != nil {
		return "-EE:EE:00"
	}

	if res.Type == timer.TimerTypeNA {
		return "-00:00:00"
	}

	if res.Type == timer.TimerTypeTime {
		return res.Time.Format(ClockFormat)
	}
	if res.Type == timer.TimerTypeDuration {
		return FormatDuration(res.Duration)
	}

	return ""
}
