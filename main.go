package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to use the simple golang CMS system.")
	http.HandleFunc("/", func(rw http.ResponseWriter, req http.Request) {

	})

}
