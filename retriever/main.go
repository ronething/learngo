package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://blog.ronething.cn")
}

func main() {
	var r Retiever

	r = mock.Retriever{"this is a fack ronething.com"}
	inspect(r)

	// Type assertion
	mockRetriever := r.(mock.Retriever)
	fmt.Println(mockRetriever.Contents)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	//fmt.Println(download(r))
}

func inspect(r Retiever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
