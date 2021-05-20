package hawk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// WebsocketSender is a Sender implementation thar sends errors via websockets.
type WebsocketSender struct {
	mx sync.Mutex
	// conn is a connection to send errors.
	conn io.ReadWriteCloser
	// addr is Hawk address
	addr string
}

// NewWebsocketSender returns new WebsocketSender instant with default address.
func NewWebsocketSender() Sender {
	return &WebsocketSender{
		mx:   sync.Mutex{},
		addr: "ws" + strings.TrimPrefix(DefaultURL, "http"),
	}
}

// connect establishes connection to Hawk.
func (w *WebsocketSender) connect() error {
	conn, _, _, err := ws.DefaultDialer.Dial(context.TODO(), w.addr)
	if err != nil {
		return err
	}
	w.conn = conn

	return nil
}

// reconnect performs attempts to connect using exponential backoff strategy.
func (w *WebsocketSender) reconnect() error {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = 2 * time.Minute
	be.InitialInterval = 100 * time.Millisecond
	be.Multiplier = 2
	be.MaxInterval = 30 * time.Second

	var err error
	b := backoff.WithContext(be, context.Background())
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return fmt.Errorf("failed to connect: %w", err)
		}
		select {
		case <-time.After(d):
			err = w.connect()
			if err != nil {
				continue
			}
			return nil
		}
	}
}

// Send sends errors to Hawk.
func (w *WebsocketSender) Send(data []ErrorReport) error {
	w.mx.Lock()
	defer func() {
		w.conn.Close()
		w.mx.Unlock()
	}()

	err := w.reconnect()
	if err != nil {
		return err
	}

	writer := wsutil.NewWriter(w.conn, ws.StateClientSide, ws.OpText)
	for _, rep := range data {
		repBytes, err := rep.MarshalJSON()
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, bytes.NewBuffer(repBytes))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

// SetURL sets addr field for WebsocketSender instance.
func (w *WebsocketSender) setURL(hawkURL string) {
	w.addr = hawkURL
}

// GetURL returns addr.
func (w *WebsocketSender) getURL() string {
	return w.addr
}
