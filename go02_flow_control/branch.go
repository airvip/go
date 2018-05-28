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

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operation:" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic("illegal score")
	case score > 90:
		g = "A"
	case score > 80:
		g = "B"
	case score > 70:
		g = "C"
	case score > 60:
		g = "D"
	default:
		g = "E"
	}
	return g
}

func main() {
	//read1()
	//read2()
	fmt.Println(eval(10, 2, "/"))
	//println(eval(10, 2, "/"))
	//println(eval(10, 2, "%"))
	fmt.Println(
		grade(45),
		grade(67),
		grade(78),
		grade(89),
		grade(100))
}
