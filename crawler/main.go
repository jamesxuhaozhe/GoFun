package main

import (
	"haozhexu.com/gofun/learngo/crawler/engine"
	"haozhexu.com/gofun/learngo/crawler/scheduler"
	"haozhexu.com/gofun/learngo/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
