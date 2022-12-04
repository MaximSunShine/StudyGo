package main

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("Свали от сюда"))
	if err != nil {
		log.Println(err)
	}
}
func InvalidMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, `NotAllowed!`)
}
func CheckStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Fine!`)
}
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `InternalServerError!`)
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `NotFound!`)
}

func main() {
	//простой способ обработки командной строки
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("hello"))
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/NotAllowed", InvalidMethod)
	http.HandleFunc("/StatusOK", CheckStatusOK)
	http.HandleFunc("/InternalServerError", InternalServerError)
	http.HandleFunc("/NotFound", NotFound)

	log.Println(http.ListenAndServe(":9090", nil))
}
