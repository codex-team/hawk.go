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
	"github.com/codex-team/hawk.go"
)

func main() {
    // initialize Hawk catcher
    catcher, err := hawk.New("abcd1234-1234-abcd-1234-123456abcdef")
    if err != nil {
        panic(err)
    }
    err = catcher.Catch(errors.New("Test exception"))
    if err != nil {
        panic(err)
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
