package main

import (
	"errors"
	"log"
	"time"

	"github.com/codex-team/hawk.go"
)

const (
	token = "abcd"
	url   = "http://server:9090"
)

var (
	ErrSomeTestError = errors.New("test error 1")
	catcher          *hawk.Catcher
	delay            = 3 * time.Second
)

func returnTestError() error {
	log.Printf("returnTestError func")
	return ErrSomeTestError
}

func test() {
	err := returnTestError()
	catcherErr := catcher.Catch(err)
	if catcherErr != nil {
		log.Fatalf("failed to catch error: %s", catcherErr.Error())
	}
}

func main() {
	log.Printf("wait %s for service availability...", delay)
	time.Sleep(delay)

	var err error
	catcher, err = hawk.New(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("set up the catcher")
	catcher.SetURL(url)
	catcher.MaxBulkSize = 1
	test()
}
