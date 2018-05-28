package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n > 0{
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
	} else if n == 0 {
		result = "0"
	}
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(0),
        convertToBin(13))
}
