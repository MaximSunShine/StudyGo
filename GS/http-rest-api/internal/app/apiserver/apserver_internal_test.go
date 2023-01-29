package apiserver

// не актуально

/*
func TestAPIServer_HandleHello(t *testing.T) {
	//создаем новый apiconfig  для тестирования, иннициализируем все поля встроенными методами структур
	s := New(NewConfig())

	//ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
	//httptest.ResponseRecorder по сути, это реализация http.ResponseWriter, которая записывает код состояния ответа,
	//заголовки и тело вместо того, чтобы фактически записывать их в HTTP-соединение. Короче говоря,
	//httptest.ResponseRecorder это реализация http.ResponseWriter,
	//которая записывает свои мутации для последующей проверки в тестах.
	rec := httptest.NewRecorder()

	// создаем спец обработчик запросов имитирующий работу сервера (фактически структура ответа сервера с необходимыми методами)
	//иннициализируем запрос к серверу (фактически структура запроса с необходимыми методами)
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	//передаем в обработчик строки запроса конкретной url: структуру ответа серверной части и структуру запроса клиентской части
	s.handleHello().ServeHTTP(rec, req)

	//assert := assert.New(t)
	//запускаем тесты на равенство ожидаемого значения ответа серверной части с готовым шаблоном/текстом
	assert.Equal(t, rec.Body.String(), "Hello")

}
*/
