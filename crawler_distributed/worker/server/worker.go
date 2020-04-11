package main

import (
	"fmt"

	"haozhexu.com/gofun/learngo/crawler/fetcher"
	"haozhexu.com/gofun/learngo/crawler_distributed/rpcsupport"
	"haozhexu.com/gofun/learngo/crawler_distributed/worker"

	"log"

	"flag"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
