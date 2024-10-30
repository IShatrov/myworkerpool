package workerpool

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
