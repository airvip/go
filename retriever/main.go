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

type Poster interface {
	Post(url string,
	    from map[string]string) string
}

const url  = "http://www.imooc.com"

func download(r Retriever) string {
    return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url,
		map[string]string{
			"name": "airvip",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}


func session(s RetrieverPoster) string {
	//s.Get()
	s.Post(url,map[string]string{
		"contents": "this is test code",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{
		"this is a fake imooc.com"}
	r = &retriever
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	// Type assertion name
	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))

	if mockRetriever,ok := r.(*mock.Retriever); ok{
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
	fmt.Println()
}