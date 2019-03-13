package atem

import "context"

func (a *ATEM) WriteEvent(ctx context.Context, toPort string, name string, data interface{}) {
	// Do not do anything (the ATEM can't receive events)
}
