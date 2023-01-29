package apiserver

import (
	"github.com/GS/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
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
}
