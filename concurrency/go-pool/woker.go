package main

// Job 任务
type Job func(id int)

// Worker work
type Worker struct {
	id   int
	quit chan bool
}

// NewWorker new worker
func NewWorker(id int) *Worker {
	return &Worker{
		id:   id,
		quit: make(chan bool),
	}
}

// Start 启动一个 Worker
func (w *Worker) Start(jobCh <-chan Job) {
	go func() {
		for {
			select {
			case job := <-jobCh:
				job(w.id)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() { w.quit <- true }()
}
