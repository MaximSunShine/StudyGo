package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandlerAll(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/NotAllowed":
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `NotAllowed!`)
	case "/StatusOK":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `Fine!`)
	case "/InternalServerError":
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `InternalServerError!`)
	case "/NotFound":
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `NotFound!`)
	}

	fmt.Fprintf(w, "Path = "+r.URL.Path)
	//fmt.Fprintf(w)
}

func RunWebServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandlerAll)

	log.Println("Запуск веб-сервера")
	err := http.ListenAndServe(":9090", mux)
	log.Fatal(err)
}

func main() {

	RunWebServer()
}
