package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hasher := md5.New()
	var b []byte
	hasher.Write([]byte("test"))
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}
