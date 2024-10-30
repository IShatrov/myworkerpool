package workerpool

import (
	"errors"
	"fmt"
	"time"
)

const bufferSize = 10


// A primitive worker-pool implementation.
type Workerpool struct {
	quitChannels map[string]chan<- bool
	data         chan string
}

// Creates a new Workerpool.
func NewWorkerpool() *Workerpool {
	quitChannels := make(map[string]chan<- bool)
	data := make(chan string, bufferSize)

	return &Workerpool{quitChannels, data}
}

// Reports whether a Workerpool is valid.
func (workerpool *Workerpool) IsValid() bool {
	return workerpool.quitChannels != nil && workerpool.data != nil
}

// Adds a worker to a Workerpool. sleepTime is used to simulate a long job.
func (workerpool *Workerpool) AddWorker(id string, sleepTime time.Duration) error {
	_, contains := workerpool.quitChannels[id]

	if contains {
		return errors.New("Pool already contains worker with id \"" + id + "\"")
	}

	quit := make(chan bool)
	workerpool.quitChannels[id] = quit

	go workerpool.worker(id, quit, workerpool.data, sleepTime)

	return nil
}

// Deletes a worker from a Workerpool.
func (workerpool *Workerpool) DeleteWorker(id string) error {
	quit, contains := workerpool.quitChannels[id]

	if contains {
		fmt.Printf("About to delete worker \"%s\"\n", id)

		select {
		case quit <- true:
			delete(workerpool.quitChannels, id)
		}

		return nil
	}

	return errors.New("Worker " + id + " does not exist")
}

// Reports whether a Workerpool has workers.
func (workerpool *Workerpool) HasWorkers() bool {
	return len(workerpool.quitChannels) != 0
}

// Adds a job to a Workerpool.
func (workerpool *Workerpool) AddJob(job string) {
	workerpool.data <- job
}

// Adds all jons in the slice to a Workerpool.
func (workerpool *Workerpool) AddJobs(jobs []string) {
	for _, job := range jobs {
		workerpool.AddJob(job)
	}
}


// Performs a job.
func (workerpool *Workerpool) worker(id string, quit <-chan bool, src <-chan string, sleepTime time.Duration) {
	for {
		select {
		case <-quit:
			fmt.Printf("Deleted worker \"%s\"\n", id)
			return
		case job := <- src:
			fmt.Printf("Worker \"%s\" about to perform job \"%s\"\n", id, job)
			time.Sleep(sleepTime)
			fmt.Printf("Worker \"%s\" finished job \"%s\"\n", id, job)
		}
	}
}
