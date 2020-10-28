package hyperdeck

import (
	"bytes"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Notification string

const (
	NotificationTransport     Notification = "transport"
	NotificationSlot          Notification = "slot"
	NotificationRemote        Notification = "remote"
	NotificationConfiguration Notification = "configuration"
)

func (c *Client) asyncRegister(writer io.Writer) string {
	uuid := uuid.New()

	c.asyncLock.Lock()
	defer c.asyncLock.Unlock()

	c.asyncWriter[uuid.String()] = writer

	return uuid.String()
}

func (c *Client) asyncUnregister(id string) {
	c.asyncLock.Lock()
	defer c.asyncLock.Unlock()

	delete(c.asyncWriter, id)
}

func (c *Client) writeAsyncPayload(payload []byte) {
	c.asyncLock.Lock()
	defer c.asyncLock.Unlock()

	c.onAsyncMessage(payload)

	for _, w := range c.asyncWriter {
		w.Write(payload)
	}
}

func (c *Client) Notify(notifs ...Notification) error {
	if len(notifs) == 0 {
		return nil
	}

	cmd := "notify:"
	for _, notif := range notifs {
		cmd = fmt.Sprintf("%s %s:true", cmd, notif)
	}

	res, err := c.Send([]byte(cmd + "\n"))
	if err != nil {
		return errors.Wrap(err, "fail to send notif packet")
	}

	if IsError(res) {
		parsedError, err := ParseError(res)
		if err != nil {
			return errors.Wrap(err, "fail to parse error")
		} else {
			return parsedError
		}
	}

	return nil
}

func (c *Client) onAsyncMessage(payload []byte) {
	if bytes.HasPrefix(payload, []byte("508 transport info")) && c.transportListener != nil {
		transport, err := ParseTransport(payload)
		if err != nil {
			// TODO cleanup
			fmt.Println("ERROR:")
			fmt.Println(err.Error())
			return
		}
		c.transportListener(transport)
	}

	if bytes.HasPrefix(payload, []byte("502 slot info")) && c.slotListener != nil {
		slot, err := ParseSlot(payload)
		if err != nil {
			// TODO cleanup
			fmt.Println("ERROR:")
			fmt.Println(err.Error())
			return
		}
		c.slotListener(slot)
	}
}
