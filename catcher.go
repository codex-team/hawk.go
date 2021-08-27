package hawk

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	// DefaultURL is the default Hawk URL to send errors.
	DefaultURL = "https://k1.hawk.so/"
	// CatcherType is the type of this Catcher.
	CatcherType = "errors/golang"
	// DefaultMaxBulkSize is default max amount of errors that can be sent at once.
	DefaultMaxBulkSize = 1
	// DefaultMaxInterval is default max time interval to wait for errors before sending them.
	DefaultMaxInterval = 5 * time.Second
	// DefaultSourceCodeLines is default number of source code lines before and
	// after the line with error that will be reported.
	DefaultSourceCodeLines = 5
)

// ErrEmptyURL is returned if an empty URL was provided in SetURL func.
var ErrEmptyURL = errors.New("empty Hawk URL")

// Catcher collects information about errors and sends them to Hawk.
type Catcher struct {
	options HawkOptions
	// lastSendTime is the time when last report was sent.
	lastSendTime time.Time
	// errorsCh is a channel where error reports will be sent.
	errorsCh chan ErrorReport
	// done is a channel that signals to stop the Catcher.
	done chan error
	// timeout
	timeout chan bool
	// sender is Sender implementation that is used by Catcher to send errors to Hawk.
	sender Sender
}

// New returns new Catcher instance with provided access token and default URL.
func New(options HawkOptions) (*Catcher, error) {
	err := checkAccessToken(options.AccessToken)
	if err != nil {
		return nil, err
	}

	// check URL
	_, err = url.Parse(options.URL)
	if err != nil {
		return nil, err
	}

	// choose and init an appropriate transport
	var sender Sender
	switch options.Transport.(type) {
	case HTTPTransport:
		sender = NewHTTPSender(options.URL, options.Debug)
	case WebsocketTransport:
		sender = NewWebsocketSender(options.URL)
	default:
		return nil, fmt.Errorf("Invalid transport value: %s", options.Transport)
	}

	catcher := &Catcher{
		options:      options,
		lastSendTime: time.Now(),
		errorsCh:     make(chan ErrorReport),
		done:         make(chan error),
		timeout:      make(chan bool, 1),
		sender:       sender,
	}

	return catcher, nil
}

// checkAccessToken validates access token.
func checkAccessToken(accessToken string) error {
	// TODO: implement token format checking
	return nil
}

// Run starts Catcher's main work to wait for errors.
func (c *Catcher) Run() error {
	var buffer []ErrorReport

	for {
		select {
		// send all errors from buffer before stop
		case err := <-c.done:
			if len(buffer) > 0 {
				sendErr := c.sender.Send(buffer)
				if sendErr != nil {
					return fmt.Errorf("failed to send errors: %s;\nCatcher exited with error: %w", sendErr, err)
				}
			}
			return err
		// process a new error message
		case report := <-c.errorsCh:
			buffer = append(buffer, report)
			if len(buffer) == c.options.MaxBulkSize {
				err := c.sender.Send(buffer)
				if err != nil {
					return fmt.Errorf("catcher exited with error: %w", err)
				}
				buffer = buffer[:0]
			} else {
				// initiate a new timer if not yet
				timer.wait(c.timeout, c.options.MaxInterval)
			}
		// send all errors from a buffer
		case <-c.timeout:
			err := c.sender.Send(buffer)
			if err != nil {
				return fmt.Errorf("catcher exited with error: %w", err)
			}
			buffer = buffer[:0]
		}
	}
}

// Stop stops Catcher.
func (c *Catcher) Stop() {
	if r := recover(); r != nil {
		c.processRecover(r)
	}
	c.done <- errors.New("stop")
}

func (c *Catcher) processRecover(r interface{}) {
	_ = c.Catch(fmt.Errorf("%s", r))
}

func (c *Catcher) Recover() {
	if r := recover(); r != nil {
		c.processRecover(r)
	}
}
