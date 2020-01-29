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

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "james", "age": "22"})
}

type RetrieverPoster interface {
	Retriever
	Poster
	// you may also declare other method
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}
	inspect(r)
	//fmt.Println(download(r))

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:" + v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:" + v.UserAgent)
	}
	fmt.Println()
}
