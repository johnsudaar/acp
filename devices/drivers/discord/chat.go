package discord

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/johnsudaar/acp/devices"
	"github.com/johnsudaar/acp/devices/types"
	"github.com/johnsudaar/acp/devices/types/chat"
	"github.com/pkg/errors"
)

type Discord struct {
	*devices.Base
	token string

	subscriptions    map[string]chat.ChatHandler
	subscriptionLock *sync.RWMutex
	discordMutex     *sync.RWMutex
	discord          *discordgo.Session
}

func (r *Discord) InputPorts() []string {
	return []string{}
}

func (r *Discord) OutputPorts() []string {
	return []string{}
}

func (r *Discord) API() http.Handler {
	return http.NotFoundHandler()
}

func (r *Discord) Start() error {
	dg, err := discordgo.New("Bot " + r.token)
	if err != nil {
		return errors.Wrap(err, "fail to create Discord dession")
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(r.onDiscordMessage)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		return errors.Wrap(err, "fail to open discord connection")
	}

	return nil

}

func (r *Discord) Stop() error {
	r.discordMutex.Lock()
	defer r.discordMutex.Unlock()
	if r.discord != nil {
		r.discord.Close()
	}

	return nil
}

func (r *Discord) WriteEvent(ctx context.Context, toPort, name string, data interface{}) {
}

func (r *Discord) WriteRealtimeEvent(ctx context.Context, channel string, payload json.RawMessage) {
}

func (r *Discord) Types() []types.Type {
	return []types.Type{
		types.ChatType,
	}
}
