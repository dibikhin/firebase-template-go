package main

import (
	"fmt"
	"log"
	"net/http"

	hello "github.com/dibikhin/firebase-template-go/internal/https"
)

func main() {
	http.HandleFunc("/", hello.HelloWorld)
	fmt.Println("Listening on http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
