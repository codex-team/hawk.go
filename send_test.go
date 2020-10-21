package hawk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var reports = []ErrorReport{
	{
		Token:       "abcd",
		CatcherType: "errors/golang",
		Payload: Payload{
			Title:     "test error 1",
			Timestamp: time.Now().String(),
			Backtrace: []Backtrace{
				{
					File: "/test/test_client.go",
					Line: 27,
					SourceCode: [SourceCodeLines]SourceCode{
						{
							LineNumber: 26,
							Content:    "\terr := returnTestError()",
						},
						{
							LineNumber: 27,
							Content:    "\tcatcherErr := catcher.Catch(err)",
						},
						{
							LineNumber: 28,
							Content:    "\tif err != nil {",
						},
					},
				},
			},
		},
	},
}

// TestSend tests sending errors by HTTP.
func TestSend(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Errorf("failed to read request body: %s", err.Error())
			}
			var res ErrorReport
			err = json.Unmarshal(bodyBytes, &res)
			if err != nil {
				t.Errorf("failed to unmarshal json body: %s", err.Error())
			}
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		c := &Catcher{
			hawkURL:   server.URL,
			client:    server.Client(),
			errBuffer: reports,
		}

		err := c.send()
		if err != nil {
			t.Errorf("failed to send errors: %s", err.Error())
		}
	})
}
