package main

// Dispatcher struct
type Dispatcher struct {
	workerpool chan Worker
	queue      chan func()
}

// NewDispatcher Create Dispatcher
func NewDispatcher(workers int, maxQueue int) Dispatcher {
	dispatcher := Dispatcher{
		workerpool: make(chan Worker, workers),
		queue:      make(chan func(), maxQueue),
	}

	for i := 0; i < workers; i++ {
		worker := NewWorker(dispatcher.workerpool)
		worker.Start()

	}

	return dispatcher
}

// AddToQueue Add to the queue
func (d Dispatcher) AddToQueue(function func()) {
	d.queue <- function
}

// Run runs the loop for the dispatcher
func (d Dispatcher) Run() {
	go d.dispatch()
}

func (d Dispatcher) dispatch() {
	for {
		select {
		case function := <-d.queue:
			go func(function func()) {
				worker := <-d.workerpool
				worker.Work(function)
			}(function)
		}
	}
}
