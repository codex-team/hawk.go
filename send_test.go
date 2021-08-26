package hawk

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var reports = []ErrorReport{
	{
		Token:       "abcd",
		CatcherType: "errors/golang",
		Payload: Payload{
			Title: "test error 1",
			Backtrace: []Backtrace{
				{
					File: "/test/test_client.go",
					Line: 27,
					SourceCode: []SourceCode{
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

// TestSendHTTP tests sending errors via HTTP client.
func TestSendHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %s", err.Error())
		}
		var res ErrorReport
		err = res.UnmarshalJSON(bodyBytes)
		if err != nil {
			t.Fatalf("failed to unmarshal json body: %s; data: %s", err.Error(), string(bodyBytes))
		}
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	t.Run("http client", func(t *testing.T) {
		s := &HTTPSender{
			addr:   server.URL,
			client: server.Client(),
		}

		err := s.Send(reports)
		if err != nil {
			t.Fatalf("failed to send errors: %s", err.Error())
		}
	})
}

// TestSendWebsockets tests sending errors via websockets.
func TestSendWebsockets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			t.Fatalf("failed to upgrade: %s", err.Error())
		}
		defer conn.Close()

		data, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			t.Fatalf("failed to read data: %s", err.Error())
		}
		var res ErrorReport
		err = res.UnmarshalJSON(data)
		if err != nil {
			t.Fatalf("failed to unmarshal json body: %s; data: %s", err.Error(), string(data))
		}
	}))

	t.Run("websockets", func(t *testing.T) {
		s := &WebsocketSender{
			addr: "ws" + strings.TrimPrefix(server.URL, "http"),
		}
		err := s.connect()
		if err != nil {
			t.Fatalf("failed to connect: %s", err.Error())
		}

		err = s.Send(reports)
		if err != nil {
			t.Errorf("failed to send errors: %s", err.Error())
		}
	})
}
