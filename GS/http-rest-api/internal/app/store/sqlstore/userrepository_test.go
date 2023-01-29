package sqlstore_test

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
	"github.com/GS/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

// добавление строки в таблицу
func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL) //создаем БД и получаем структуру и фн удаление таблиц с бд
	defer teardown("users")                         // удаление таблицы

	s := sqlstore.NewStore(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u)) // проверка ошибки
	assert.NotNil(t, u)                   // проверка наличия заполненной структуры
}

// Поиск строки по емэйлу
func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL) //создаем БД и получаем структуру и фн удаление таблиц с бд
	defer teardown("users")                         // удаление таблицы

	s := sqlstore.NewStore(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email) //Поиск строки по емэйлу
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t) //создаем тестового пользователя
	u.Email = email
	s.User().Create(u) //добавляем в его в таблицу

	u, err = s.User().FindByEmail(email) //находим пользователя по емэйлу и возвращаем
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
