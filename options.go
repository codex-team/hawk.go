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
	// URL is the Hawk endpoint
	URL string
	// Release is the custom application version
	Release string
}

func DefaultHawkOptions() HawkOptions {
	return HawkOptions{
		MaxBulkSize:       DefaultMaxBulkSize,
		MaxInterval:       DefaultMaxInterval,
		SourceCodeEnabled: false,
		SourceCodeLines:   DefaultSourceCodeLines,
		AccessToken:       "",
		URL:               "",
	}
}

type HawkAdditionalParams func(payload *Payload)
