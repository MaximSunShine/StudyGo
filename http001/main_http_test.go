package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned %v", status)
	}
	expect := `Fine!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}
func TestInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/NotAllowed", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InvalidMethod)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned %v", status)
	}
	expect := `NotAllowed!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}
func TestInternalServerError(t *testing.T) {
	req, err := http.NewRequest("GET", "/InternalServerError", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InternalServerError)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusInternalServerError {
		t.Errorf("handler returned %v", status)
	}
	expect := `InternalServerError!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}
func TestNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/NotFound", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFound)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned %v", status)
	}
	expect := `NotFound!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}
