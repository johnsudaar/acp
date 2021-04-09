package discord

import (
	"context"

	"github.com/johnsudaar/acp/devices/types/chat"
)

var _ chat.Chattable = &Discord{}

func (r *Discord) ChatSubscribe(ctx context.Context, cb chat.ChatHandler) string {
	return ""
}

func (r *Discord) ChatUnsubscribe(ctx context.Context, id string) {

}
