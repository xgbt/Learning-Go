package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf.
	// fmt.Fprintf() 依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了 io.Writer 接口。
	// Fprintf() 能够写入任何类型，只要其实现了 Write 方法，包括 os.Stdout,文件（例如 os.File），管道，网络连接，通道等等，
	// 同样的也可以使用 bufio 包中缓冲写入。
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()

}
