package function

import (
	"net/http"

	hello "github.com/dibikhin/firebase-template-go/internal/https"
)

// Hello again
func Hello(w http.ResponseWriter, r *http.Request) {
	hello.HelloWorld(w, r)
}
