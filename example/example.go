package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/codex-team/hawk.go"
)

func main() {
	options := hawk.DefaultHawkOptions()
	options.AccessToken = "<TOKEN>"
	options.Domain = "stage-k1.hawk.so"
	options.Debug = true
	options.Transport = hawk.HTTPTransport{}
	options.AffectedUser = hawk.AffectedUser{Id: "01", Name: "default user"}

	catcher, err := hawk.New(options)
	if err != nil {
		log.Fatal(err)
	}

	go catcher.Run()
	defer catcher.Stop()

	err = catcher.Catch(fmt.Errorf("manual exception without context"))
	if err != nil {
		catcher.Stop()
		log.Fatal(err)
	}

	err = catcher.Catch(fmt.Errorf("manual exception with context"),
		hawk.WithContext(struct{ Timestamp string }{Timestamp: strconv.Itoa(int(time.Now().Unix()))}),
		hawk.WithUser(hawk.AffectedUser{Id: "uid", Name: "N0str"}),
		hawk.WithRelease("v-3.7"),
	)
	if err != nil {
		catcher.Stop()
		log.Fatal(err)
	}

	// panic
	var s []interface{}
	fmt.Println(s[10])

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
