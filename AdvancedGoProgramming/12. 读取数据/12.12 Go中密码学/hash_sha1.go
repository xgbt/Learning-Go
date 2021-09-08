package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)
var b []byte

func main() {
	// 通过调用 sha1.New() 创建了一个新的 hash.Hash 对象，用来计算 SHA1 校验值。
	// Hash 类型实际上是一个接口，它实现了 io.Writer 接口：
	hasher := sha1.New()
	// 通过 io.WriteString 或 hasher.Write 将给定的 []byte 附加到当前的 hash.Hash 对象中。
	io.WriteString(hasher, "test")


	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	// hasher.Write(data) 更加直观
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checkSum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checkSum)
}
