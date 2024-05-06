package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hi there, I love %s", request.URL.Path)
}

func viewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]

	page, error := loadPage(title)

	if error != nil {
		fmt.Println("Error while loading page")
	}

	fmt.Fprintf(responseWriter, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func allHandlers(responseWriter http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[1:]

	fmt.Println(path)

	if strings.Contains(path, "view/") {
		viewHandler(responseWriter, request)
	} else {
		handler(responseWriter, request)
	}
}

func main() {
	http.HandleFunc("/", allHandlers)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
