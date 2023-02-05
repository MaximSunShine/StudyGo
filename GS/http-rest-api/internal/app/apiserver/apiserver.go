package apiserver

import (
	"database/sql"
	"github.com/GS/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

// Start подключение к БД, и запуск роутера со всеми оставшимися конфигурациями
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL) //подключение к БД и тест/пинг ее работы
	if err != nil {
		return err
	}

	defer db.Close() // закрыть подключение после завершения работы роутера
	store := sqlstore.NewStore(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore) // создаем структуру для работы с БД

	return http.ListenAndServe(config.BindAddr, srv) // запуск сервера
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL) // подключаемся к БД и возвращаем структуру БД
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	} // проверка подключения к БД на плучение строки

	return db, nil
}

/*
// APIServer объявлеие структуры
type APIServer struct {
	config *Config        //(struct)
	logger *logrus.Logger //(struct)объявляем поле типа структура Logger с большим кол параметров и методов для логирования
	router *mux.Router    // объявляем роутер
	store  *sqlstore.Store
}

// New создаем переменную типа структура и возвращаем ее параметром возврата фн
func New(cf *Config) *APIServer {
	return &APIServer{
		config: cf,
		logger: logrus.New(),    //иннициализироем структуру Logger встроенным методом
		router: mux.NewRouter(), //иннициализируем структуру Router со встроенными методами
	}
}

// Start главный метод начала работы сервера , настройка конфигурации
func (aps *APIServer) Start() error {

	if err := aps.configureLogger(); err != nil {
		return err
	} //конфигурируем логгер
	aps.configureRouter() // конфигурируем роутер

	if err := aps.configureStore(); err != nil {
		return err
	}

	aps.logger.Info("starting api server") //записываем в лог текстовую информацию
	// запускаем сервер на прослушивание урл адресов с нужного порта
	return http.ListenAndServe(aps.config.BindAddr, aps.router)
}

// configureLogger настраиваем логгер
func (aps *APIServer) configureLogger() error {
*/
/*	func ParseLevel(lvl string) (Level, error) {

	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	}
*/
/*//ParseLevel
	level, err := logrus.ParseLevel(aps.config.LogLevel) // считваем код типа обработки ("error","debug")?
	if err != nil {
		return err
	}
	// устанавливаем на какой тип просушка логер будет вестись, меньше число, выше приоритет см выше?
	aps.logger.SetLevel(level)
	return nil
}

// configureRouter настраиваем роутер
func (aps *APIServer) configureRouter() {
	//привязваем конкретный урл адрес к определенному адресу
	aps.router.HandleFunc("/hello", aps.handleHello())
}

func (aps *APIServer) configureStore() error {
	st := sqlstore.NewStore(aps.config.Store)
	if err := st.Open(); err != nil {
		aps.logger.Info(err) //
		return err
	}
	aps.store = st
	return nil

}

// handleHello проводим замыкание фн хэндлера (ServeHTTP) с
// возможностью обратиться и изменить необходимые параметры в теле данной фн
func (aps *APIServer) handleHello() http.HandlerFunc {
	//замыкание
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello")
		if err != nil {
			return
		}
	}
}
*/
