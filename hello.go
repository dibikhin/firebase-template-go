package hello

import (
	"fmt"
	"net/http"
	"time"
)

// https://<region>-<project>.cloudfunctions.net/HelloWorld
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey ddude!", time.Now())
}
