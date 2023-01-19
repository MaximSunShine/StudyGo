package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	//Помощник помечает вызывающую функцию как тестовую вспомогательную функцию.
	//При печати информации о файле и строке эта функция будет пропущена.
	//Помощник может быть вызван одновременно из нескольких подпрограмм.
	t.Helper()

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
}
