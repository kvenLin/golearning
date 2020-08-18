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

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "clf",
		TimeOut:   time.Minute,
	}
	inspect(r)

	fmt.Println("Type assertion: ")
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
}

func inspect(r Retriever) {
	fmt.Println("Type Switch: ")
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content: ", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
