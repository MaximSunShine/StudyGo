package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL") //читаем поле DATABASE_URL, и оно скорее всего пустое
	if databaseURL == "" {
		//"postgres://postgres:domcrat@localhost:5432/restapi_dev?sslmode=disable"
		databaseURL = "host=localhost user=postgres password=domcrat dbname=restapi_test  sslmode=disable" //
	}
	os.Exit(m.Run()) // запуск тестирования
}
