package switcher

import "github.com/johnsudaar/acp/devices/types"

func Import(device interface{}) (types.DeviceType, error) {
	switchable, ok := device.(Switchable)
	if !ok {
		return nil, types.ErrInvalidType
	}
	driver := NewSwitcherDriver(switchable)
	return driver, nil
}
