package main

import (
	"fmt"
	"go/retriever/mock"
	"go/retriever/real"
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
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	// Type assertion name
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))
}

func inspect(r Retriever)  {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}