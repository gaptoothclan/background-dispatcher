package main

type Worker struct {
	workerpool chan Worker
	function   chan func()
	quit       chan bool
}

func NewWorker(workerpool chan Worker) Worker {
	return Worker{
		workerpool: workerpool,
		function:   make(chan func()),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerpool <- w
			select {
			case function := <-w.function:
				function()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Work(function func()) {
	w.function <- function
}

func (w Worker) Quit() {
	w.quit <- true
}
