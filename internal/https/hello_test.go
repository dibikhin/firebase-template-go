package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	HelloWorld(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Hello StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	}
}
