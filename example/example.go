package main

import (
	"errors"
	"fmt"
	"github.com/codex-team/hawk.go"
	"sync"
)

func main() {
	catcher, err := hawk.InitWithUrl("token", "http://localhost:3000/catcher")
	if err != nil {
		panic(err)
	}
	_ = catcher.Catch(errors.New("Test exception"))

	parallelTest(catcher)
}

func parallelTest(catcher *hawk.Catcher)  {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			_ = catcher.Catch(errors.New(fmt.Sprintf("Test exception №%d", i)))
			wg.Done()
		}(i)
	}
	wg.Wait()
}