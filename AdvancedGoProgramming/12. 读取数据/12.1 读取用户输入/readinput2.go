package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个读取器，并将其与标准输入绑定。
	// bufio.NewReader() 构造函数的签名为: func NewReader(rd io.Reader) *Reader
	// 该函数的实参可以是满足 io.Reader 接口的任意对象（任意包含有适当的 Read() 方法的对象)
	// 函数返回一个新的带缓冲的 io.Reader 对象，它将从指定读取器（例如 os.Stdin）读取内容。
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some input: ")

	if input, err := inputReader.ReadString('\n'); err == nil {
		fmt.Println("The input was: ", input)
	}
}