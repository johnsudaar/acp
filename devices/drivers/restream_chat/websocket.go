package restream

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Scalingo/go-utils/logger"
	"github.com/gorilla/websocket"
)

type RestreamPayload struct {
	Action    string          `json:"action,omitempty"`
	Payload   json.RawMessage `json:"payload,omitempty"`
	Timestamp uint64          `json:"timestamp,omitempty"`
}

type RestreamEvent struct {
	ConnectionIdentifier string               `json:"connectionIdentifier,omitempty"`
	EventIdentifier      string               `json:"eventIdentifier,omitempty"`
	EventSourceID        int                  `json:"eventSourceId,omitempty"`
	EventTypeID          int                  `json:"eventTypeId,omitempty"`
	UserID               int                  `json:"userId,omitempty"`
	EventPayload         RestreamEventPayload `json:"eventPayload,omitempty"`
}

type RestreamEventPayload struct {
	Author RestreamEventAuthor `json:"author,omitempty"`
	Bot    bool                `json:"bot,omitempty"`
	Text   string              `json:"text,omitempty"`
}

type RestreamEventAuthor struct {
	Avatar      string `json:"avatar,omitempty"`
	Color       string `json:"color,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
}

const backendURL = "wss://backend.chat.restream.io/ws/embed"

func (r *Restream) openWebsocketConnection(ctx context.Context) {
	log := logger.Get(ctx)
	url := fmt.Sprintf("%s?token=%s", backendURL, r.token)

	var payload RestreamPayload
	var chatEvent RestreamEvent
	for {
		if r.isStopped() {
			return
		}

		time.Sleep(1 * time.Second)
		log.Info("Connecting...")
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.WithError(err).Error("fail to connect to restream ws")
			break
		}
		r.connMutex.Lock()
		r.conn = c
		r.connMutex.Unlock()
		log.Info("Connected")

		for {
			if r.isStopped() {
				return
			}

			c.SetReadDeadline(time.Now().Add(30 * time.Second))
			_, message, err := c.ReadMessage()
			if err != nil {
				log.WithError(err).Error("fail to receive message")
				break
			}
			err = json.Unmarshal(message, &payload)
			if err != nil {
				log.WithError(err).Error("fail to decode event")
				continue
			}
			if payload.Action != "event" {
				continue
			}
			err = json.Unmarshal(payload.Payload, &chatEvent)
			if err != nil {
				log.WithError(err).Error("fail to decode payload")
				continue
			}

			if chatEvent.EventPayload.Bot || chatEvent.EventPayload.Text == "" {
				continue
			}

			log.Infof("[%s] %s", chatEvent.EventPayload.Author.DisplayName, chatEvent.EventPayload.Text)
		}

	}

}
