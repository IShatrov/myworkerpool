package main

import (
	"time"

	"github.com/IShatrov/myworkerpool/workerpool"
)

func main() {
	pool := workerpool.NewWorkerpool()

	pool.AddWorker("Amogus", time.Second)
	pool.AddWorker("Sus", time.Second)

	pool.DeleteWorker("Amogus")

	pool.AddJob("Skjfnd")
	pool.AddJob("Sdlmkfdn")

	time.Sleep(5 * time.Second)
}
