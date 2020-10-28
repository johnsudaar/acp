package hyperdeck

import (
	"bytes"
	"context"
	"io"
	"net"
	"strings"
	"time"

	"github.com/Scalingo/go-utils/logger"
	"github.com/pkg/errors"
)

type Operation struct {
	Payload     []byte
	Result      chan []byte
	ResultError chan error
}

func (c *Client) Send(payload []byte) ([]byte, error) {
	operation := Operation{
		Payload:     payload,
		Result:      make(chan []byte),
		ResultError: make(chan error),
	}
	c.operations <- operation

	select {
	case payload := <-operation.Result:
		return payload, nil
	case err := <-operation.ResultError:
		return nil, errors.Wrap(err, "fail to execute error")
	}
}

func (c *Client) writer() {
	for {
		c.stopLock.Lock()
		stopping := c.stopping
		c.stopLock.Unlock()
		if stopping {
			return
		}

		<-c.writeSync
		operation := <-c.operations
		c.operationSync.Lock()
		c.curOperation = &operation
		c.operationSync.Unlock()

		c.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := c.conn.Write(operation.Payload)
		if err != nil {
			if err == io.EOF {
				c.Stop()
				continue
			}
			c.operation().ResultError <- errors.Wrap(err, "fail to send message")
			c.resetOperation()
			go func() {
				c.writeSync <- true
			}()
			continue
		}
	}
}

func (c *Client) reader(ctx context.Context) {
	log := logger.Get(ctx)
	buff := make([]byte, 1024)
	c.writeSync <- true
	curBuff := make([]byte, 0)

	for {
		c.stopLock.Lock()
		stopping := c.stopping
		c.stopLock.Unlock()
		if stopping {
			return
		}

		c.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := c.conn.Read(buff)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				c.Stop()
				continue
			}
			if err == io.EOF {
				c.Stop()
				continue
			}
			if c.operation() == nil {
				log.WithError(err).Error("Fail to read (unsolicited)")
				continue
			} else {
				c.operation().ResultError <- errors.Wrap(err, "fail to read response")
				c.resetOperation()
				continue
			}
		}

		curBuff = append(curBuff, buff[:n]...)
		if n >= 150 && !bytes.HasSuffix(curBuff, []byte{0x0d, 0x0a, 0x0d, 0x0a}) {
			continue
		}
		res := make([]byte, len(curBuff))
		copy(res, curBuff)
		curBuff = make([]byte, 0)
		if strings.HasPrefix(string(res), "5") {
			c.writeAsyncPayload(res)
			continue
		}

		if c.operation() == nil {
			log.WithField("msg", string(res)).Error("Unsolicited message from hyperdeck")
			continue
		}

		c.operation().Result <- res
		c.resetOperation()
		c.writeSync <- true
	}
}

func (c *Client) operation() *Operation {
	c.operationSync.Lock()
	defer c.operationSync.Unlock()
	return c.curOperation
}

func (c *Client) resetOperation() {
	c.operationSync.Lock()
	defer c.operationSync.Unlock()
	c.curOperation = nil
}
