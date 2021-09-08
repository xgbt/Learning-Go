package main

import "os"

func main() {
	// 使用 os.Stdout.WriteString("hello, world\n")，我们可以输出到屏幕。
	os.Stdout.WriteString("hello, world\n")
	file, _ := os.OpenFile("goprogram", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	// 我们不使用缓冲区，直接将内容写入文件：f.WriteString()
	file.WriteString("hello, world\n")
}
