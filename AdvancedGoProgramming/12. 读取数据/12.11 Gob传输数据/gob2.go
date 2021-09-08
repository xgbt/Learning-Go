package main

import (
	"encoding/gob"
	"log"
	"os"
)


var content string

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("encode error:", err)
	}


}
