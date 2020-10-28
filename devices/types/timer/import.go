package timer

import "github.com/johnsudaar/acp/devices/types"

func Import(device interface{}) (types.DeviceType, error) {
	timeable, ok := device.(Timeable)
	if !ok {
		return nil, types.ErrInvalidType
	}

	driver := NewTimerDriver(timeable)
	return driver, nil
}
