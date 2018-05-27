package main

import (
	"io/ioutil"
	"fmt"
)

func read1(){
	const filename  = "E:/phpStudy/WWW/go/go02_flow_control/abc.txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", b)
	}
}

func read2()  {
	const filename  = "E:/phpStudy/WWW/go/go02_flow_control/abc.txt"
	if b, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", b)
	}
}

func main() {
	read1()
	read2()
}
