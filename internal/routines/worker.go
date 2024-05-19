package routines

import (
	"counter/internal/models"
	"sync"
)

func Worker(wg *sync.WaitGroup, jobs <-chan models.Numbers, counter *models.Counter, done chan struct{}) {
	defer wg.Done()
	for {
		select {
		case numbers, ok := <-jobs:
			if !ok {
				return
			}
			sum := numbers.A + numbers.B
			counter.Mu.Lock()
			counter.TotalSum += sum
			counter.Mu.Unlock()
		case <-done:
			return
		}
	}
}
