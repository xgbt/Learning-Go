package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 扩展 cat.go 例子，使用 flag 添加一个选项，目的是为每一行头部加入一个行号。使用 cat -n goprogram 测试输出。

var numberFlag = flag.Bool("n", false, "number each line")

func cat2(r *bufio.Reader) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}

	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat2(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat2(bufio.NewReader(f))
		f.Close()
	}
}