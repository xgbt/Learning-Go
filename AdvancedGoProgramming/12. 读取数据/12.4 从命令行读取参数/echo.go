package main

import (
	"flag"
	"os"
)

// echo.go 模拟了 Unix 的 echo 功能

// NewLine echo -n flag, of type *bool
// flag.Bool() 定义了一个默认值是 false 的 flag：
// 当在命令行出现了第一个参数（这里是 "n"），flag 被设置成 true（NewLine 是 *bool 类型）。
// flag 被解引用到 *NewLine，所以当值是 true 时将添加一个 Newline（"\n"）。
var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	// flag.PrintDefaults() 打印 flag 的使用帮助信息，本例中打印的是:
	// -n=false: print newline
	flag.PrintDefaults()
	// flag.Parse() 扫描参数列表（或者常量列表）并设置 flag
	// Parse() 之后 flag.Arg(i) 全部可用，flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字。
	flag.Parse()
	var s = ""
	// flag.Narg() 返回参数的数量
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			// -n is parsed, flag becomes true
			if *NewLine {
				s += Newline
			}
		}
		// flag.Arg(i) 表示第i个参数
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}
