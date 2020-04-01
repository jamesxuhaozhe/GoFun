package main

import (
	"haozhexu.com/gofun/learngo/crawler/engine"
	"haozhexu.com/gofun/learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

}
