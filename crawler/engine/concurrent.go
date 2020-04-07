package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDupURL(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			if isDupURL(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedURLs = make(map[string]bool)

func isDupURL(url string) bool {
	if visitedURLs[url] {
		return true
	}

	visitedURLs[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
