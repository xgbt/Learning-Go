package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	bks := make([]Book, 1)

	filename := "D:\\Workspace\\go\\Learning-Go\\AdvancedGoProgramming\\12. 读取数据\\12.2 文件读写\\products.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error %s opening %s: ", err, filename)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file
		line, err := reader.ReadString('\n')
		readErr := err
		// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
		line = string(line[:len(line)-2])
		fmt.Printf("The input was: -%s-\n", line)

		book := new(Book)
		strSl := strings.Split(line, ";")
		// read title
		book.title = strSl[0]
		// read price
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		// read quantity
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		// 这里不能直接append(bks, *book), 因为数组里默认有bks[0]
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
