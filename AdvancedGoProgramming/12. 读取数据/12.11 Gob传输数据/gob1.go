package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// 一个编解码，并且以字节缓冲模拟网络传输的简单例子：
type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)

}
