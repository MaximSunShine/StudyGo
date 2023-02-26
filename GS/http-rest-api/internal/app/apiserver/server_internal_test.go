package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name:         "authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}
	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

// сверка вводимых данных логин пароль
func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@exmple.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	//сверка ожидаемых данных при передаче входных данных с возвращаемыми
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

// сверка уже хранимых данных логин пароль
func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	//сверка ожидаемых данных при передаче входных данных с возвращаемыми
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

/*
	//ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
	//httptest.ResponseRecorder по сути, это реализация http.ResponseWriter, которая записывает код состояния ответа,
	//заголовки и тело вместо того, чтобы фактически записывать их в HTTP-соединение. Короче говоря,
	//httptest.ResponseRecorder это реализация http.ResponseWriter,
	//которая записывает свои мутации для последующей проверки в тестах.
	// Cоздаем спец обработчик запросов имитирующий работу сервера
	//(фактически структура ответа сервера с необходимыми методами)
	rec := httptest.NewRecorder() //ResponseWriter

	// создаем запрос от клиента к серверу
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	//создаем новый тестовый сервер с привязкой всех ручек
	s := newServer(teststore.New())

	//передаем в обработчик строку запроса конкретной url: структуру ответа серверной части
	//и структуру запроса клиентской части
	s.ServeHTTP(rec, req) // обрабатываем запрос на стороне сервера и возвращаем ответ (в данном случае создаем иммитацию)

	//запускаем тесты на равенство ожидаемого значения ответа серверной части с готовым шаблоном/текстом
	assert.Equal(t, rec.Code, http.StatusOK)

*/
