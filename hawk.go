package hawk

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultURL is the default Hawk URL to send errors.
	DefaultURL = "https://hawk.so/catcher/golang"
	// CatcherType is the type of this Catcher.
	CatcherType = "errors/golang"
	// DefaultMaxBulkSize is default max amount of errors that can be sent at once.
	DefaultMaxBulkSize = 64
	// DefaultMaxInterval is default max time interval to wait for errors before sending them.
	DefaultMaxInterval = 5 * time.Minute
)

// ErrEmptyURL is returned if an empty URL was provided in SetURL func.
var ErrEmptyURL = errors.New("empty Hawk URL")

// Catcher collects information about errors and sends them to Hawk.
type Catcher struct {
	// hawkURL is the address to send errors.
	hawkURL string
	// accessToken is the Hawk access token.
	accessToken string
	// maxBulkSize is max amount of errors that can be sent at once.
	MaxBulkSize int
	// maxInterval is max time interval to wait for errors before sending them.
	MaxInterval time.Duration
	// lastSendTime is the time when last report was sent.
	lastSendTime time.Time
	// errBuffer stores last N errors for MaxInterval time, N <= MaxBulkSize.
	errBuffer []ErrorReport
	// HTTP client
	client *http.Client
}

// New returns new Catcher instance with provided access token and default URL.
func New(token string) (*Catcher, error) {
	err := checkAccessToken(token)
	if err != nil {
		return nil, err
	}

	catcher := &Catcher{
		hawkURL:      DefaultURL,
		accessToken:  token,
		MaxBulkSize:  DefaultMaxBulkSize,
		MaxInterval:  DefaultMaxInterval,
		lastSendTime: time.Now(),
		errBuffer:    []ErrorReport{},
		client:       &http.Client{},
	}

	return catcher, nil
}

// checkAccessToken validates access token.
func checkAccessToken(accessToken string) error {
	return nil
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
	c.hawkURL = hawkURL

	return nil
}

// clearBuffer resets stored errors.
func (c *Catcher) clearBuffer() {
	c.errBuffer = c.errBuffer[:0]
}
