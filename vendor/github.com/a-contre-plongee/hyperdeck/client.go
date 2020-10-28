package hyperdeck

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	addr string
	conn net.Conn

	repeaterAddr string
	repeaterConn net.Listener

	motd []byte

	operations    chan Operation
	curOperation  *Operation
	operationSync *sync.Mutex
	writeSync     chan bool

	asyncWriter map[string]io.Writer
	asyncLock   *sync.Mutex

	stopping bool
	stopLock *sync.Mutex

	transportListener TransportListener
	slotListener      SlotListener
}

type ClientOpts func(c *Client)

func WithRepeater(addr string) ClientOpts {
	return func(c *Client) {
		c.repeaterAddr = addr
	}
}

func WithTransportListener(listener TransportListener) ClientOpts {
	return func(c *Client) {
		c.transportListener = listener
	}
}

func WithSlotListener(listener SlotListener) ClientOpts {
	return func(c *Client) {
		c.slotListener = listener
	}
}

func New(addr string, opts ...ClientOpts) *Client {
	client := &Client{
		addr:          addr,
		operations:    make(chan Operation),
		operationSync: &sync.Mutex{},
		writeSync:     make(chan bool),
		asyncWriter:   make(map[string]io.Writer),
		asyncLock:     &sync.Mutex{},
		stopLock:      &sync.Mutex{},
		stopping:      false,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *Client) Start(ctx context.Context) error {
	c.stopping = false
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:9993", c.addr), 1*time.Second)
	if err != nil {
		return errors.Wrap(err, "fail to open connection")
	}

	// Wait for Motd
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	buff := make([]byte, 4096)
	n, err := conn.Read(buff)
	if err != nil {
		conn.Close()
		return errors.Wrap(err, "fail to test connection (read)")
	}
	conn.SetReadDeadline(time.Time{})
	c.motd = buff[:n]
	c.conn = conn

	go c.reader(ctx)
	go c.writer()
	go c.watchdog()

	err = c.Notify(NotificationTransport, NotificationSlot)
	if err != nil {
		panic(err)
	}

	// We are succesfully connected
	if c.repeaterAddr != "" {
		err := c.repeater(ctx)
		if err != nil {
			c.Stop()
			return errors.Wrap(err, "fail to start repeater")
		}
	}

	return nil
}

func (c *Client) Running() bool {
	c.stopLock.Lock()
	defer c.stopLock.Unlock()
	return !c.stopping
}

func (c *Client) Stop() {
	c.stopLock.Lock()
	c.stopping = true
	c.stopLock.Unlock()

	if c.conn != nil {
		c.conn.Close()
	}
	if c.repeaterConn != nil {
		c.repeaterConn.Close()
	}
}

func (c *Client) watchdog() {
	for {
		c.stopLock.Lock()
		stopping := c.stopping
		c.stopLock.Unlock()
		if stopping {
			return
		}

		c.Send([]byte("ping\n"))
		time.Sleep(8 * time.Second)
	}
}
