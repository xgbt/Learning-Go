package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 我们以只写模式打开文件 output.dat，如果文件不存在则自动创建：
	// 可以看到，OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符“|”连接），使用的文件权限。
	//
	// 我们通常会用到以下标志：
	// os.O_RDONLY：只读
	// os.O_WRONLY：只写
	// os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	// os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
	//
	// 在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用0。
	// 而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creating\n")
		return
	}
	defer outputFile.Close()

	// 创建一个写入器（缓冲区）对象
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello world!\n"

	// 将字符串写入缓冲区，写 10 次
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}

	// 缓冲区的内容紧接着被完全写入文件
	outputWriter.Flush()
}
