package apiserver

import (
	"github.com/GS/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router    // роутер
	logger *logrus.Logger // логирование
	store  store.Store    // база данных
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter() // приязываем хэндлеры/ручки к обработчику роутера

	return s
}

// метод необходимый для реализации интерфейса
//
//	type Handler interface {
//		ServeHTTP(ResponseWriter, *Request)
//	}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r) //запускает метод роутера ServeHTTP
}

func (s *server) configureRouter() {
	//привязываем обработчик к конктерному урл с указанным методом запроса
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")

}

// обработчик/ручка
func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
