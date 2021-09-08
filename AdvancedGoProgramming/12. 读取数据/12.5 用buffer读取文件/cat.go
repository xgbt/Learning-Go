package main

// 结合使用了缓冲读取文件和命令行 flag 解析这两项技术。如果不加参数，那么你输入什么屏幕就打印什么。
// 参数被认为是文件名，如果文件存在的话就打印文件内容到屏幕。命令行执行 cat goprogram 测试输出。
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("No Args!")
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.OpenFile(flag.Arg(i), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
