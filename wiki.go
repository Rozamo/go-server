package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"

	return os.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, error := os.ReadFile((filename))

	if error != nil {

		return nil, error
	}

	return &Page{title, body}, nil
}

func exMain() {
	page1 := &Page{Title: "Test Page", Body: []byte("This is a sample page")}
	page1.save()

	page2, error := loadPage("Test Page")

	if error != nil {
		fmt.Println("Error while loading page")
	}

	fmt.Println(string(page2.Body))
}
