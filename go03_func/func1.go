package main

import "fmt"

func swap(a, b int)  {
	a, b = b, a
}

// 指针
func swap1(a, b *int)  {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b) // a == 3, b == 4

	swap1(&a, &b)
	fmt.Println(a, b) // a == 4, b == 3

	a, b = swap2(a, b)
    fmt.Println(a, b)
}
