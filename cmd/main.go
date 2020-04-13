package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dibikhin/firebase-template-go"
)

func main() {
	http.HandleFunc("/", hello.HelloWorld)
	fmt.Println("listening on localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
