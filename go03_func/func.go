package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		panic("illegal operation:" + op)
	}
}

func div(a, b int) (int, int) {
    return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(eval(20, 5, "/"))
	fmt.Println(div(13, 5))
	q, r := div2(13, 5)
	fmt.Println(q, r)
}
