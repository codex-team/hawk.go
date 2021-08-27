# Hawk catcher for Golang

Golang errors catcher for [Hawk.so](https://hawk.so)

## Installation

### Go Get

```golang
go get https://github.com/codex-team/hawk.go
```

## Usage

Initiate Hawk catcher
```
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
```

Run Hawk errors collection and panic recovery
``` 
go catcher.Run()
defer catcher.Stop()
```

You can manually send errors and context information
```
err = catcher.Catch(fmt.Errorf("manual exception with context"),
    hawk.WithContext(struct{ Timestamp string }{Timestamp: strconv.Itoa(int(time.Now().Unix()))}),
    hawk.WithUser(hawk.AffectedUser{Id: "uid", Name: "N0str"}),
    hawk.WithRelease("v-3.7"),
)
```

Full code example

```golang
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

```

# About CodeX

<img align="right" width="120" height="120" src="https://codex.so/public/app/img/codex-logo.svg" hspace="50">

CodeX is a team of digital specialists around the world interested in building high-quality open source products on a global market. We are [open](https://codex.so/join) for young people who want to constantly improve their skills and grow professionally with experiments in cutting-edge technologies.

| 🌐 | Join  👋  | Twitter | Instagram |
| -- | -- | -- | -- |
| [codex.so](https://codex.so) | [codex.so/join](https://codex.so/join) |[@codex_team](http://twitter.com/codex_team) | [@codex_team](http://instagram.com/codex_team) |

## License

MIT
