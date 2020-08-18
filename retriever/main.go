package main

import (
	"fmt"
	"imooc.com/louye/learngo/retriever/mock"
	"imooc.com/louye/learngo/retriever/real"
	"time"
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
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPost interface {
	Retriever
	Poster
}

func session(s RetrieverPost) string {
	s.Post(url, map[string]string{
		"content": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Content: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "clf",
		TimeOut:   time.Minute,
	}
	inspect(r)

	fmt.Println("Type assertion: ")
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println("Try Session: ")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting:", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type Switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content: ", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}
