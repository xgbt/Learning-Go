package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	chars int
	words int
	lines int
)

func Counters(input string) {
	chars += len(input) - 2 // -2 for \r\n
	words += len(strings.Fields(input))
	lines++
}

func main() {
	chars, words, lines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Eneter some input, type S to stop: ")

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred: ", err)
		}
		if input == "S\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", chars)
			fmt.Printf("Number of words: %d\n", words)
			fmt.Printf("Number of lines: %d\n", lines)
			os.Exit(0)
		}
		Counters(input)
	}
}
