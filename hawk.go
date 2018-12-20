package hawk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const defaultUrl = "https://hawk.so/catcher/golang"

type Catcher struct {
	url string
	AccessToken AccessToken
}

type AccessToken struct {
	token string
}

func (catcher *Catcher) SetEndpoint (hawkUrl string) error {
	if hawkUrl == "" {
		return errors.New("empty catcher URL")
	}

	uri, err := url.Parse(hawkUrl)
	if err != nil {
		return err
	}

	catcher.url = uri.String()
	return nil
}

func New(accessToken AccessToken) (*Catcher, error) {
	catcher := &Catcher{defaultUrl, AccessToken{""}}
	err := checkAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	catcher.AccessToken = accessToken
	return catcher, nil
}

func Init(accessToken string) (*Catcher, error) {
	catcher, err := New(AccessToken{accessToken})
	if err != nil {
		return nil, err
	}
	return catcher, nil
}

func InitWithUrl(accessToken string, hawkUrl string) (*Catcher, error) {
	catcher, err := New(AccessToken{accessToken})
	if err != nil {
		return catcher, err
	}
	err = catcher.SetEndpoint(hawkUrl)
	if err != nil {
		return catcher, err
	}
	return catcher, nil
}

func checkAccessToken(accessToken AccessToken) error {
	return nil
}

func (catcher *Catcher) catch(error_data ErrorData) error {
	client := &http.Client{
	}

	error_bytes, err := json.Marshal(error_data)
	if err != nil {
		return err
	}

	message := Request{catcher.AccessToken.token, error_bytes, "errors/golang", Sender{"127.0.0.1"}}
	message_bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = client.Post(catcher.url, "application/json", bytes.NewBuffer(message_bytes))

	return err
}

func (catcher *Catcher) CatchWithCode(error_data error, code int) error {
	return catcher.catch(ErrorData{code, fmt.Sprintf("%s", error_data)})
}

func (catcher *Catcher) Catch(error_data error) error {
	return catcher.catch(ErrorData{0, fmt.Sprintf("%s", error_data)})
}