package store_test

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

// добавление строки в таблицу
func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL) //создаем БД и получаем структуру и фн удаление таблиц с бд
	defer teardown("users")                        // удаление таблицы

	u, err := s.User().Create(&model.User{ // добавление строки в таблицу
		Email: "user@example.org",
	})
	assert.NoError(t, err) // проверка ошибки
	assert.NotNil(t, u)    // проверка наличия заполненной структуры
}

// Поиск строки по емэйлу
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL) //создаем БД и получаем структуру и фн удаление таблиц с бд
	defer teardown("users")                        // удаление таблицы

	email := "user@example.org"
	u, err := s.User().FindByEmail(email) //Поиск строки по емэйлу
	assert.Error(t, err)

	u, err = s.User().Create(&model.User{ //создаем струку с искомым емэйлом и проверяем ее поиск
		Email: "user@example.org",
	})
	u, err = s.User().FindByEmail(email) //поиск должен завершиться удачно
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
