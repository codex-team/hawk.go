package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/codex-team/hawk.go"
)

func main() {
	catcher, err := hawk.New("token", hawk.NewHTTPSender())
	if err != nil {
		log.Fatal(err)
	}

	err = catcher.SetURL("http://localhost:3000/catcher")
	if err != nil {
		log.Fatal(err)
	}

	err = catcher.Catch(fmt.Errorf("Test exception"))
	if err != nil {
		log.Fatal(err)
	}

	parallelTest(catcher)
}

func parallelTest(catcher *hawk.Catcher) {
	go catcher.Run()
	defer catcher.Stop()
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			err := catcher.Catch(fmt.Errorf("Test exception №%d", i))
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
