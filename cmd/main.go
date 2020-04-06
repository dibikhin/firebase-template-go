package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/hello"
)

func main() {
	http.HandleFunc("/", hello.HelloWorld)
	fmt.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
