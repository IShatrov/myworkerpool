package main

import (
	"time"

	"github.com/IShatrov/myworkerpool/workerpool"
)

func main() {
	pool := workerpool.NewWorkerpool()

	pool.AddJob("Emergency meeting")
	pool.AddJob("Vote red")

	pool.AddWorker("Amogus", time.Second)
	pool.AddWorker("Impostor", time.Second)

	pool.AddJob("Eject red")

	time.Sleep(5 * time.Second)
}
