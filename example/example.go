package main

import (
	"time"

	"github.com/IShatrov/myworkerpool/workerpool"
)

func main() {
	pool := workerpool.NewWorkerpool()

	pool.AddJobs([]string{"Job One", "Job Two", "Job Three"})

	pool.AddWorker("Worker One", time.Second)
	pool.AddWorker("Worker Two", time.Second)

	pool.DeleteWorker("Worker Two")

	pool.AddJob("Job Four")

	time.Sleep(5 * time.Second) // Naive way to wait until all jobs are done
}
