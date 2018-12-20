package hawk

import "encoding/json"

type ErrorData struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
type Sender struct {
	Ip string `json:"ip"`
}

type Request struct {
	Token string `json:"token"`
	Payload json.RawMessage `json:"payload"`
	CatcherType string `json:"catcher_type"`
	Sender Sender `json:"sender"`
}