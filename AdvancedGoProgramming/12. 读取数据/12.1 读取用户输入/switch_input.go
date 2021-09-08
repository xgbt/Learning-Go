package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")

	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Your name is %s", input)
	} else {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	// For Unix: goprogram with delimiter "\n", for Windows: goprogram with "\r\n"
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("U are not welcome here !\n")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Chris\r\n":
		fallthrough
	case "Ivo\r\n":
		fmt.Printf("Welcome %s!\n", input)
	default:
		fmt.Printf("U are not welcome here !\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Chris\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("U are not welcome here !\n")
	}
}
