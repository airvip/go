package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename  = "E:/phpStudy/WWW/go/go02_flow_control/abc.txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", b)
	}
}
