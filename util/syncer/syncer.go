package syncer

import "sync"

var wg sync.WaitGroup

func SetAlive() {
	wg.Add(1)
}

func SetCompleted() {
	wg.Done()
}

func Wait() {
	wg.Wait()
}
