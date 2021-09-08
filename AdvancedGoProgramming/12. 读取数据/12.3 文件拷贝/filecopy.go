package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.OpenFile(srcName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	copyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}
