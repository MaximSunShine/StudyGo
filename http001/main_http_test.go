package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type request struct {
	url    string
	status int
}

func TestHandlerAll(t *testing.T) {

	var r = []request{
		{"/NotAllowed", http.StatusMethodNotAllowed},
		{"/StatusOK", http.StatusOK},
		{"/InternalServerError", http.StatusInternalServerError},
		{"/NotFound", http.StatusNotFound}}
	for _, val := range r {
		req, err := http.NewRequest("GET", val.url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HandlerAll)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != val.status {
			t.Errorf("handler returned %v", status)
		}
		/*expect := `Fine!`
		if rr.Body.String() != expect {
			t.Errorf("handler returned %v", rr.Body.String())
		}*/
	}
}
