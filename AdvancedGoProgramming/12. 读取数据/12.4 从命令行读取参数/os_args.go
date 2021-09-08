package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice"
	// 这个命令行参数会放置在切片 os.Args[] 中（以空格分隔），从索引1开始（os.Args[0] 放的是程序本身的名字，在本例中是 os_args）。
	// 函数 strings.Join 以空格为间隔连接这些参数。
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
