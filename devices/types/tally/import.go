package tally

import (
	"github.com/johnsudaar/acp/devices/types"
)

func Import(device interface{}) (types.DeviceType, error) {
	tallyable, ok := device.(Tallyable)
	if !ok {
		return nil, types.ErrInvalidType
	}

	driver := NewTallyDriver(tallyable)
	return driver, nil
}
