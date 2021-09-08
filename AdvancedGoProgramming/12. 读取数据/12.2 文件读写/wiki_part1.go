package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() (err error) {
	return ioutil.WriteFile(page.Title, page.Body, 0666)
}

func (page *Page) load(title string) (err error) {
	page.Title = title
	page.Body, err = ioutil.ReadFile(title)
	return
}

func main() {
	page := Page{"Page.md", []byte("# Page\n## Section1\nThis is section1.")}
	page.save()

	var newPage Page
	newPage.load("Page.md")
	fmt.Println(string(newPage.Body))
}
