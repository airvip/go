package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

// for 初始值 条件 变化量
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

// 省略初始值，条件
func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	/*scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}*/

	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 省略初始值，条件，变化量
func forever() {
	for { // 死循环
		fmt.Println("airvip")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(0),
        convertToBin(13))

	printFile("E:/phpStudy/WWW/go/src/go/go02_flow_control/abc.txt")

	s := `abc"d"
    kkk
    123
    
    airivp`

    printFileContents(strings.NewReader(s))
	//forever()
}
