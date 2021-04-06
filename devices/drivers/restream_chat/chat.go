package restream

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/chat"
)

type Restream struct {
	*devices.Base
	token string

	subscriptions    map[string]chat.ChatHandler
	subscriptionLock *sync.RWMutex

	connMutex *sync.Mutex
	conn      *websocket.Conn

	stop      bool
	stopMutex *sync.RWMutex
}

func (r *Restream) InputPorts() []string {
	return []string{}
}

func (r *Restream) OutputPorts() []string {
	return []string{}
}

func (r *Restream) API() http.Handler {
	return http.NotFoundHandler()
}

func (r *Restream) Start() error {
	go r.openWebsocketConnection(context.Background())
	return nil
}

func (r *Restream) Stop() error {
	r.stopMutex.Lock()
	r.stop = true
	r.stopMutex.Unlock()

	r.connMutex.Lock()
	if r.conn != nil {
		r.conn.Close()
	}
	r.connMutex.Unlock()
	return nil
}

func (r *Restream) isStopped() bool {
	r.stopMutex.RLock()
	defer r.stopMutex.RUnlock()
	return r.stop
}

func (r *Restream) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (r *Restream) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}

func (r *Restream) Types() []types.Type {
	return []types.Type{
		types.ChatType,
	}
}
