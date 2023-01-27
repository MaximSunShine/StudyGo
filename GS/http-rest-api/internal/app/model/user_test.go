package model_test

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {

	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		}, //valid
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u

			},
			isValid: false,
		}, //empty email
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u

			},
			isValid: false,
		}, //invalid email
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u

			},
			isValid: false,
		}, //empty password
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"

				return u

			},
			isValid: false,
		}, //short password
		{
			name: "encrypt password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.EncryptedPassword = "encriptedpasword"
				u.Password = ""

				return u

			},
			isValid: true,
		}, //with encrypt password
	}

	//перебор всех вариантов структуры
	for _, tc := range testCases {
		//тестовый run с переданной функцией проверяет: будет ли ошибка при неверных данных
		//и наоборот (суб тесты, запускаются параллельно)
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())     // проверка успешной генерации хэш-пароля
	assert.NotEmpty(t, u.EncryptedPassword) // проверка наличия хэшпароля
}
