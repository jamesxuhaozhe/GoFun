package scheduler

import "haozhexu.com/gofun/learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	go func() { s.workerChan = c }()
}
