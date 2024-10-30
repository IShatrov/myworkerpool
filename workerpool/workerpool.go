package workerpool

import (
	"errors"
	"fmt"
	"time"
)

const bufferSize = 10

type Workerpool struct {
	quitChannels map[string]chan<- bool
	data         chan string
}

func NewWorkerpool() *Workerpool {
	quitChannels := make(map[string]chan<- bool)
	data := make(chan string, bufferSize)

	return &Workerpool{quitChannels, data}
}

func (workerpool *Workerpool) IsValid() bool {
	return workerpool.quitChannels != nil && workerpool.data != nil
}

func (workerpool *Workerpool) AddWorker(id string, sleepTime time.Duration) error {
	_, contains := workerpool.quitChannels[id]

	if contains {
		return errors.New("Pool already contains worker with id" + id)
	}

	quit := make(chan bool)
	workerpool.quitChannels[id] = quit

	go workerpool.worker(id, quit, workerpool.data, sleepTime)

	return nil
}

func (workerpool *Workerpool) AddJob(job string) {
	workerpool.data <- job
}

func (workerpool *Workerpool) worker(id string, quit <-chan bool, src <-chan string, sleepTime time.Duration) {
	for {
		select {
		case <-quit:
			fmt.Println(id)
			return
		default:
			job := <-src
			fmt.Printf("Worker \"%s\" about to perform job \"%s\"\n", id, job)
			time.Sleep(sleepTime)
			fmt.Printf("Worker \"%s\" finished job \"%s\"\n", id, job)
		}
	}
}
