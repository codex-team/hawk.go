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
	options.MaxInterval = 1 * time.Second
	options.AccessToken = "<TOKEN>"
	options.URL = "https://test.stage-k1.hawk.so"
	options.Release = "v1.3.3"

	catcher, err := hawk.New(options, hawk.NewHTTPSender())
	if err != nil {
		log.Fatal(err)
	}

	go catcher.Run()
	defer catcher.Stop()

	err = catcher.Catch(fmt.Errorf("exception NEW"),
		hawk.WithContext(struct{ Timestamp string }{Timestamp: strconv.Itoa(int(time.Now().Unix()))}),
		hawk.WithUser(hawk.AffectedUser{Id: "uid", Name: "N0str"}),
		hawk.WithRelease("v-3.7"),
	)
	if err != nil {
		catcher.Stop()
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
