package threading

import (
	"github.com/google/uuid"
	"time"
)

type Thread struct {
	Name     string
	Callable func()
	id       string
	isAlive  bool
	end      chan bool
}

func (t *Thread) Id() string {
	return t.id
}

func (t *Thread) Start() {
	t.isAlive = true
	t.end = make(chan bool, 1)
	t.id = uuid.New().String()

	go t.run()
}

// Join timeOut - is time to wait for the end of the goroutine execution in seconds
func (t *Thread) Join(timeOut int) {
	if timeOut > 0 {
		for t.isAlive && timeOut > 0 {
			time.Sleep(time.Second)
			timeOut--
		}
	} else {
		<-t.end
	}
}

func (t *Thread) run() {
	t.Callable()

	t.isAlive = false
	t.end <- true
}

func (t *Thread) IsAlive() bool {
	return t.isAlive
}
