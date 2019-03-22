package routinepool

import (
	"log"
)

type WorkOp func(data interface{}) interface{}
type WorkNotice func(data interface{})

func NewWork(data interface{}, loop WorkOp, notice WorkNotice) {
	routinePool.jobQueue <- work{
		data:   data,
		loop:   loop,
		notice: notice,
	}
}

type work struct {
	data   interface{}
	loop   WorkOp
	notice WorkNotice
}

type worker struct {
	WorkPool   chan chan work
	jobChannel chan work
	quit       chan bool
}

func NewWorker(workerPool chan chan work) worker {
	return worker{
		WorkPool:   workerPool,
		jobChannel: make(chan work),
		quit:       make(chan bool),
	}
}
func (w worker) Working() {
	go func() {
		for {
			w.WorkPool <- w.jobChannel
			select {
			case job := <-w.jobChannel:
				job.notice(job.loop(job.data))
			case <-w.quit:
				if len(w.jobChannel) > 0 {
					log.Println("Working not complete!")
				}
				return
			}
		}
	}()
}
func (w worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
