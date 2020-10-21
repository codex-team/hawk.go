package hawk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// send sends errors to Hawk.
func (c *Catcher) send() error {
	for _, rep := range c.errBuffer {
		reqBytes, err := json.Marshal(rep)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPost, c.hawkURL, bytes.NewBuffer(reqBytes))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp, err := c.client.Do(req)
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

// proceedReport calls send func if there is MaxBulkSize errors caught or
// MaxInterval exceeded and clears errBuffer.
func (c *Catcher) proceedReport(report *ErrorReport) error {
	c.errBuffer = append(c.errBuffer, *report)
	if len(c.errBuffer) == c.MaxBulkSize {
		err := c.send()
		if err != nil {
			return err
		}
		c.lastSendTime = time.Now()
		c.clearBuffer()
	} else if c.lastSendTime.Add(c.MaxInterval).Before(time.Now()) {
		err := c.send()
		if err != nil {
			return err
		}
		c.lastSendTime = time.Now()
	}

	return nil
}
