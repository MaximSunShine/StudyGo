package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

var count int

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//fmt.Println(count)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(count)))

		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		if i, err := strconv.Atoi(r.FormValue("count")); err == nil {
			count += i
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.Itoa(count)))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))

		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Метод " + r.Method + " не реализован!"))

}
func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
