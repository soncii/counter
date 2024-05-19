package app

import (
	"counter/internal/models"
	"counter/internal/routines"
	"counter/internal/util"
	"fmt"
	"sync"
)

func Run(numWorkers int, filePath string) {
	numbers, err := util.ReadJSONFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	results := models.Counter{}
	var wg sync.WaitGroup
	done := make(chan struct{})

	go func() {
		results = processNumbers(numbers, numWorkers, &wg, done)
		close(done)
	}()

	Shutdown(&wg, done)

	fmt.Printf("Total sum: %d\n", results.TotalSum)
}

func processNumbers(numbers []models.Numbers, numWorkers int, wg *sync.WaitGroup, done chan struct{}) models.Counter {
	numJobs := len(numbers)
	jobs := make(chan models.Numbers, numJobs)
	results := models.Counter{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go routines.Worker(wg, jobs, &results, done)
	}

	for _, number := range numbers {
		select {
		case jobs <- number:
		case <-done:
			break
		}
	}
	close(jobs)
	wg.Wait()
	return results
}
