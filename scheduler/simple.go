package scheduler

import "github.com/hulunbao/golang/learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}

func (s SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
