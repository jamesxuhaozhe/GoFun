package main

import (
	"haozhexu.com/gofun/learngo/crawler/engine"
	"haozhexu.com/gofun/learngo/crawler/persist"
	"haozhexu.com/gofun/learngo/crawler/scheduler"
	"haozhexu.com/gofun/learngo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
