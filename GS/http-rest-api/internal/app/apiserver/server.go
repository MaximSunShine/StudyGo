package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	sessionNmae = "gopherschool"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *mux.Router // роутер
	//logger *logrus.Logger // логирование
	store        store.Store // база данных
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		//logger: logrus.New(),
		store:        store,
		sessionStore: sessionStore,
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
	s.router.HandleFunc("/sessions", s.handleSessionCreate()).Methods("POST")
}

// обработчик/ручка
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct { // объявляем структуру для заполнения при декодирования данных с бади ответа
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//замыкание
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{} //создаем новую структуру
		//получаем данные с бади и заполняем структуру
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// если не удалось декодировать запускаем метод обработки ошибок
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitaze()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct { // объявляем структуру для заполнения при декодирования данных с бади ответа
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{} //создаем новую структуру
		//получаем данные с бади и заполняем структуру
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// если не удалось декодировать запускаем метод обработки ошибок
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		/*type Session struct {
			ID      string
			Values  map[interface{}]interface{}
			Options *Options
			IsNew   bool
			store   Store
			name    string
		}*/
		//сохраняем сессию по работе каждого пользователя
		session, err := s.sessionStore.Get(r, sessionNmae)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	//приводим тип ошибки к виду хэша, который далее джейсон закодирует
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code) // записываем в заголовок код статуса ошибки
	if data != nil {
		json.NewEncoder(w).Encode(data) //выводим название ошибки (есть пользовательское именование)
	}
}
