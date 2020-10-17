package atem

import (
	"context"
	"encoding/json"
)

func (a *ATEM) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	// Do not do anything (the ATEM can't receive events)
}

func (a *ATEM) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
	// Do not do anything (the ATEM can't receive realtime events)
}
