package hawk

import (
	"errors"
	"log"
	"net/url"
	"time"
)

const (
	// DefaultURL is the default Hawk URL to send errors.
	DefaultURL = "https://k1.hawk.so/"
	// CatcherType is the type of this Catcher.
	CatcherType = "errors/golang"
	// DefaultMaxBulkSize is default max amount of errors that can be sent at once.
	DefaultMaxBulkSize = 64
	// DefaultMaxInterval is default max time interval to wait for errors before sending them.
	DefaultMaxInterval = 5 * time.Minute
	// DefaultSourceCodeLines is default number of source code lines before and
	// after the line with error that will be reported.
	DefaultSourceCodeLines = 5
)

// ErrEmptyURL is returned if an empty URL was provided in SetURL func.
var ErrEmptyURL = errors.New("empty Hawk URL")

// Sender is used to perform network interaction between Catcher and Hawk.
type Sender interface {
	// Send sends error reports to Hawk.
	Send([]ErrorReport) error
	// setURL is used to set custom URL.
	setURL(string)
	// GetURL returns Sender's URL.
	getURL() string
}

// Catcher collects information about errors and sends them to Hawk.
type Catcher struct {
	// maxBulkSize is max amount of errors that can be sent at once.
	MaxBulkSize int
	// maxInterval is max time interval to wait for errors before sending them.
	MaxInterval time.Duration
	// SourceCodeEnabled sets whether source code is available and should be
	// collected in backtrace.
	SourceCodeEnabled bool
	// SourceCodeLines is number of source code lines before and
	// after the line with error that will be reported.
	SourceCodeLines int

	// accessToken is the Hawk access token.
	accessToken string
	// lastSendTime is the time when last report was sent.
	lastSendTime time.Time
	// errorsCh is a channel where error reports will be sent.
	errorsCh chan ErrorReport
	// done is a channel that signals to stop the Catcher.
	done chan error
	// sender is Sender implementation that is used by Catcher to send errors to Hawk.
	sender Sender
}

// New returns new Catcher instance with provided access token and default URL.
func New(token string, s Sender) (*Catcher, error) {
	err := checkAccessToken(token)
	if err != nil {
		return nil, err
	}

	catcher := &Catcher{
		accessToken:       token,
		MaxBulkSize:       DefaultMaxBulkSize,
		MaxInterval:       DefaultMaxInterval,
		SourceCodeEnabled: true,
		SourceCodeLines:   DefaultSourceCodeLines,
		lastSendTime:      time.Now(),
		errorsCh:          make(chan ErrorReport),
		done:              make(chan error),
		sender:            s,
	}

	return catcher, nil
}

// checkAccessToken validates access token.
func checkAccessToken(accessToken string) error {
	// TODO: implement JWT token validation
	return nil
}

// Run starts Catcher's main work to wait for errors.
func (c *Catcher) Run() {
	buffer := []ErrorReport{}
	for {
		select {
		case err := <-c.done:
			if err != nil {
				log.Fatal(err)
			}
			return
		case report := <-c.errorsCh:
			buffer = append(buffer, report)
			if len(buffer) == c.MaxBulkSize {
				err := c.sender.Send(buffer)
				if err != nil {
					c.done <- err
					break
				}
				c.lastSendTime = time.Now()
				buffer = buffer[:0]
			} else if c.lastSendTime.Add(c.MaxInterval).Before(time.Now()) && (len(buffer) != 0) {
				err := c.sender.Send(buffer)
				if err != nil {
					c.done <- err
					break
				}
				c.lastSendTime = time.Now()
				buffer = buffer[:0]
			}
		}
	}
}

// Stop stops Catcher.
func (c *Catcher) Stop() {
	close(c.done)
}

// SetURL sets hawkURL field for Catcher instance.
func (c *Catcher) SetURL(hawkURL string) error {
	if hawkURL == "" {
		return ErrEmptyURL
	}

	_, err := url.Parse(hawkURL)
	if err != nil {
		return err
	}
	c.sender.setURL(hawkURL)

	return nil
}

// GetURL returns Sender's URL.
func (c *Catcher) GetURL() string {
	return c.sender.getURL()
}
