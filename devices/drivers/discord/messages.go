package discord

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/johnsudaar/acp/devices/types/chat"
	"github.com/johnsudaar/acp/events"
)

func (d *Discord) onDiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Type != discordgo.MessageTypeDefault {
		return
	}

	if m.Content == "" {
		return
	}

	timestamp, err := m.Timestamp.Parse()
	if err != nil {
		timestamp = time.Now()
	}

	content, err := m.ContentWithMoreMentionsReplaced(s)
	if err != nil {
		content = m.Content
	}

	msg := chat.ChatMessage{
		ID:        m.ID,
		Timestamp: timestamp.Unix(),
		From:      m.Author.Username,
		Message:   content,
		Channel:   "N/C",
	}
	d.PublishRealtimeEvent(context.TODO(), events.ChatChannel, msg)
}
