package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 0
var bb = "airvip"
var cc = true

var (
	name = "airvip"
	age = 25
	country = "china"
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter()  {
	a, b, c, d := 3, 4, true, "abc"
	b = 10
	fmt.Println(a, b, c, d)
}

func euler()  {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Println(
		cmplx.Exp(1i * math.Pi) + 1)
		//cmplx.Pow(math.E, 1i * math.Pi) + 1)

	fmt.Printf("%.3f \n",
		cmplx.Exp(1i * math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts(){
	const filename  = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums()  {
	const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

func enums1()  {
	const(
		cpp = iota
        _
		python
		golang
		javascript
	)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, name, age, country)

	euler()
	triangle()
	consts()
	enums()
	enums1()
}
