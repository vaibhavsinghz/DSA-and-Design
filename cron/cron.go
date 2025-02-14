package main

import (
	"fmt"
	"sync"
	"time"
)

func scheduleJob(runAt time.Time, job func(), wg *sync.WaitGroup) {
	duration := time.Until(runAt)
	if duration < 0 {
		fmt.Println("Specified time is in the past. Cannot schedule job.")
		return
	}

	fmt.Printf("Job scheduled to run at: %s\n", runAt)
	time.AfterFunc(duration, func() {
		defer wg.Done() // Mark the job as done when it completes
		job()
	})
}

func main() {
	// Define multiple jobs
	jobs := []struct {
		name   string
		runAt  time.Time
		action func()
	}{
		{"Job 1", time.Now().Add(5 * time.Second), func() { fmt.Println("Job 1 executed at:", time.Now()) }},
		{"Job 2", time.Now().Add(10 * time.Second), func() { fmt.Println("Job 2 executed at:", time.Now()) }},
		{"Job 3", time.Now().Add(15 * time.Second), func() { fmt.Println("Job 3 executed at:", time.Now()) }},
	}

	var wg sync.WaitGroup

	// Schedule each job
	for _, job := range jobs {
		wg.Add(1) // Increment the counter for each job
		scheduleJob(job.runAt, job.action, &wg)
	}

	// Wait for all jobs to complete
	fmt.Println("Waiting for all jobs to execute...")
	wg.Wait()
	fmt.Println("All jobs executed.")
	select {}
}
