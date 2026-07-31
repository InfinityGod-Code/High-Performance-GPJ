package channel

import (
	"fmt"
	"sync"
)

// ====================================================================
// WORKER POOL — fan-out jobs to multiple workers, fan-in results
// ====================================================================

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2 // simulate work
	}
}

func WorkerPoolDemo() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Launch workers (consumers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs (producer)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // signals workers that no more jobs are coming

	// Wait for all workers to finish, then close results
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
