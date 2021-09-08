package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	// 使用 bufio.NewReader 来获得一个读取器变量
	inputReader := bufio.NewReader(inputFile)
	for {
		// 在使用 ReadString 和 ReadBytes 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了。
		// 另外，我们也可以使用 ReadLine() 方法来实现相同的功能。
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		// 一旦读取到文件末尾，变量 readerError 的值将变成非空（事实上，其值为常量 io.EOF）
		if readerError == io.EOF {
			return
		}
	}
}
