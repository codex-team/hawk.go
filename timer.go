package hawk

import (
	"time"

	"golang.org/x/sync/semaphore"
)

type Timer struct {
	sem *semaphore.Weighted
}

var timer = Timer{sem: semaphore.NewWeighted(1)}

func (t *Timer) wait(ch chan bool, interval time.Duration) {
	if !t.sem.TryAcquire(1) {
		return
	}

	go func() {
		defer t.sem.Release(1)
		time.Sleep(interval)
		ch <- true
	}()
}
