package teststore_test

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
	"github.com/GS/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

// добавление строки в таблицу
func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()                  //создаем пустую инниц. подобие бд (мапа)
	u := model.TestUser(t)                // создаем модель структуры пользователя
	assert.NoError(t, s.User().Create(u)) // создаем пользователя и проверяем на ошибку
	assert.NotNil(t, u)                   // проверка наличия заполненной структуры
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t) //создаем тестового пользователя
	s.User().Create(u1)     //добавляем в его в таблицу

	u2, err := s.User().Find(u1.ID) //находим пользователя по емэйлу и возвращаем
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

// Поиск строки по емэйлу
func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email) //Поиск строки по емэйлу
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t) //создаем тестового пользователя
	u.Email = email
	s.User().Create(u) //добавляем в его в таблицу/мапу, нет проверки на ошибку

	u, err = s.User().FindByEmail(email) //находим пользователя по емэйлу и возвращаем
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
