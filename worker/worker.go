package main

import (
	"sync"
)

type WorkerPool struct {
	TaskChannel chan Task
	Concurrency int
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		TaskChannel: make(chan Task),
		Concurrency: concurrency,
	}
}

func (wp *WorkerPool) Worker() {
	for task := range wp.TaskChannel {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.Concurrency; i++ {
		go wp.Worker()
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.wg.Add(1)
	wp.TaskChannel <- task

}

func (wp *WorkerPool) Stop() {
	close(wp.TaskChannel)
	wp.wg.Wait()
}
