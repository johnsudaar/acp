package chat

import "github.com/johnsudaar/acp/devices/types"

func Import(device interface{}) (types.DeviceType, error) {
	chattable, ok := device.(Chattable)
	if !ok {
		return nil, types.ErrInvalidType
	}
	driver := NewChatDriver(chattable)
	return driver, nil
}
