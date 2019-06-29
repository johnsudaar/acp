package ptz

import "github.com/johnsudaar/acp/devices/types"

func Import(device interface{}) (types.DeviceType, error) {
	ptzable, ok := device.(Ptzable)
	if !ok {
		return nil, types.ErrInvalidType
	}
	driver := NewPtzDriver(ptzable)
	return driver, nil
}
