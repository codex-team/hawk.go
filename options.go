package hawk

import (
	"time"
)

// HawkOptions is used to setup Hawk Catcher
type HawkOptions struct {
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
	AccessToken string
	// Domain is the Hawk base endpoint
	Domain string
	// Release is the custom application version
	Release string
	// Whether to log debug messages
	Debug bool
	// Transport for error sending: HTTPTransport, WebsocketTransport
	Transport Transport
	// Global affected user
	AffectedUser AffectedUser
	// Full URL of the Hawk endpoint
	FullURL string
}

func DefaultHawkOptions() HawkOptions {
	return HawkOptions{
		MaxBulkSize:       DefaultMaxBulkSize,
		MaxInterval:       DefaultMaxInterval,
		SourceCodeEnabled: false,
		SourceCodeLines:   DefaultSourceCodeLines,
		AccessToken:       "",
		Domain:            "k1.hawk.so",
		Debug:             false,
		Transport:         HTTPTransport{},
		AffectedUser:      AffectedUser{},
	}
}

type HawkAdditionalParams func(payload *Payload)
