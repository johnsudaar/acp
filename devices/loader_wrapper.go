package devices

import (
	"context"
	"encoding/json"

	"github.com/johnsudaar/acp/devices/params"
)

type LoaderWrapper struct {
	Implementation DeviceLoader
}

func (l *LoaderWrapper) Load(ctx context.Context, base *Base, params json.RawMessage) (Device, error) {
	device, err := l.Implementation.Load(ctx, base, params)
	if err != nil {
		return nil, err
	}

	return Wrap(device)
}

func (l *LoaderWrapper) Validate(msg json.RawMessage) error {
	return l.Implementation.Validate(msg)
}

func (l *LoaderWrapper) Params() params.Params {
	return l.Implementation.Params()
}

func WrapLoader(l DeviceLoader) DeviceLoader {
	return &LoaderWrapper{
		Implementation: l,
	}
}
