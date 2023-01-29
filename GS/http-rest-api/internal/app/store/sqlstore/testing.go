package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// Создание тестовой БД
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	//Помощник помечает вызывающую функцию как тестовую вспомогательную функцию.
	//При печати информации о файле и строке эта функция будет пропущена.
	//Помощник может быть вызван одновременно из нескольких подпрограмм.
	t.Helper()

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	//возвращаем БД и фн удаление всех таблиц БД после тестов
	return db, func(tables ...string) {
		if len(tables) > 0 {
			//db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
	}

	/*
		config := NewConfig()            // создаем новую ConfigStore структуру
		config.DatabaseUrl = databaseURL // привязываем к полю бд новое подключение (логин пароль и тд)
		s := NewStore(config)            // создаем новую структуру Store и заполняем значениями из конфига
		//проверям бд на подключение и пингуем ее
		if err := s.Open(); err != nil {
			t.Fatal(err)
		}
		//возвращаем Store структуру и фн которая удалит все таблицы после тестов, очищая ее для последующих тестов
		return s, func(tables ...string) {
			if len(tables) > 0 {
				//удаление всех таблиц
				if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
					t.Fatal(err)
				}
			}
			s.Close() // закрытие подключения к бд
		}

	*/
}
