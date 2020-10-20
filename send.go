package hawk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// send sends errors to Hawk.
func (c *Catcher) send() error {
	repBytes, err := json.Marshal(c.errBuffer)
	if err != nil {
		return err
	}
	resp, err := c.client.Post(c.hawkURL, "application/json", bytes.NewBuffer(repBytes))
	if err != nil {
		return err

	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
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
