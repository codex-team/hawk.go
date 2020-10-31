# Hawk catcher for Golang

Golang errors catcher for [Hawk.so](https://hawk.so)

## Installation

### Go Get

```golang
go get https://github.com/codex-team/hawk.go
```

## Usage

```golang
package main

import(
  "fmt"
  "log"

	"github.com/codex-team/hawk.go"
)

func main() {
    // initialize Hawk Catcher
    catcher, err := hawk.New("abcd1234-1234-abcd-1234-123456abcdef", hawk.NewHTTPSender())
    if err != nil {
        log.Fatal(err)
    }

    go catcher.Run()
    defer catcher.Stop()

    err = catcher.Catch(fmt.Errorf("Test exception"))
    if err != nil {
        log.Fatal(err)
    }
}
```

## Issues and improvements

Feel free to ask questions or improve the project.

## Links

Repository: https://github.com/codex-team/hawk.go

Report a bug: https://github.com/codex-team/hawk.go/issues

CodeX Team: https://ifmo.su

## License

MIT
