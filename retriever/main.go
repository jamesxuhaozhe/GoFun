package main

import (
	"fmt"
	"time"

	"haozhexu.com/gofun/learngo/retriever/mock"
	"haozhexu.com/gofun/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}
	inspect(r)
	//fmt.Println(download(r))

	// type assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:" + v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:" + v.UserAgent)
	}
}
