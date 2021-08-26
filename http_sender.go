package hawk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPSender is a Sender implementation thar sends errors via http.Client.
type HTTPSender struct {
	// addr is Hawk address
	addr string
	// client is HTTP client
	client *http.Client
}

// NewHTTPSender returns new HTTPSender instant with default address.
func NewHTTPSender() Sender {
	return &HTTPSender{
		addr:   DefaultURL,
		client: &http.Client{},
	}
}

// Send sends errors to Hawk.
func (h *HTTPSender) Send(data []ErrorReport) error {
	for _, rep := range data {
		reqBytes, err := rep.MarshalJSON()
		if err != nil {
			return err
		}

		log.Printf("%s\n", reqBytes)

		req, err := http.NewRequest(http.MethodPost, h.addr, bytes.NewBuffer(reqBytes))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp, err := h.client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send errors: %w", err)
		}
		defer resp.Body.Close()
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("\n\tstatus code: %d,\n\t body: %s\n\t payload: %s", resp.StatusCode, string(respBytes), string(reqBytes))
		}
	}

	return nil
}

// SetURL sets addr field for setURL instance.
func (h *HTTPSender) setURL(hawkURL string) {
	h.addr = hawkURL
}

// GetURL returns addr.
func (h *HTTPSender) getURL() string {
	return h.addr
}
