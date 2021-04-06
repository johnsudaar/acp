package restream

import (
	"context"

	"github.com/johnsudaar/acp/devices/types/chat"
)

var _ chat.Chattable = &Restream{}

func (r *Restream) ChatSubscribe(ctx context.Context, cb chat.ChatHandler) string {
	return ""
}

func (r *Restream) ChatUnsubscribe(ctx context.Context, id string) {

}
