package threading

type Worker struct {
	ThreadCount int
	JoinThreads bool
	Callable    func()
	threads     []*Thread
}

func (w *Worker) Run() {
	if w.ThreadCount < 1 {
		w.ThreadCount = 1
	}

	w.startThreads()

	if w.JoinThreads {
		w.joinThreads()
	}
}

func (w *Worker) joinThreads() {
	for _, thread := range w.threads {
		thread.Join(0)
	}
}

func (w *Worker) startThreads() {
	i := 0
	for i < w.ThreadCount {
		w.startNewThread()
		i++
	}
}

func (w *Worker) startNewThread() {
	t := &Thread{
		Callable: w.Callable,
	}
	t.Start()
	w.threads = append(w.threads, t)
}
