package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	filename := "MyFile.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], filename, err)
		os.Exit(1)
	}
	defer file.Close()

	var reader *bufio.Reader
	filezipped, err := gzip.NewReader(file)
	if err != nil {
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(filezipped)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
