package hyperdeck

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/Scalingo/go-utils/logger"
	"github.com/pkg/errors"
)

func (c *Client) repeater(ctx context.Context) error {
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:9993", c.repeaterAddr))
	if err != nil {
		return errors.Wrap(err, "fail to listen addr")
	}

	c.repeaterConn = conn
	go c.repeaterAcceptLoop(ctx)
	return nil
}

func (c *Client) repeaterAcceptLoop(ctx context.Context) {
	log := logger.Get(ctx)
	for {
		c.stopLock.Lock()
		stopping := c.stopping
		c.stopLock.Unlock()
		if stopping {
			return
		}

		conn, err := c.repeaterConn.Accept()
		if err != nil {
			log.WithError(err).Error("Fail to accept connection")
			continue
		}

		go c.handleConnection(ctx, conn)
	}
}

func (c *Client) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	log := logger.Get(ctx)
	conn.Write(c.motd)
	subscriberId := c.asyncRegister(conn)
	defer c.asyncUnregister(subscriberId)

	buff := make([]byte, 1024*1024)

	for {
		c.stopLock.Lock()
		stopping := c.stopping
		c.stopLock.Unlock()
		if stopping {
			return
		}

		n, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				log.Info("Client closed connection")
				return
			}
			log.WithError(err).Error("Fail to read client")
			return
		}

		res, err := c.Send(buff[:n])
		if err != nil {
			log.WithError(err).Error("Fail to send command")
			return
		}

		_, err = conn.Write(res)
		if err != nil {
			log.WithError(err).Error("fail to write")
			return
		}
	}
}
