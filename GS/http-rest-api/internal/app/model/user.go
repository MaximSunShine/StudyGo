package model

import (
	val "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct { // создаем аналогию таблицы базы данных
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate проверка на соответсвие заполнения емэйла и пароля по общепринятым соглашениям
func (u *User) Validate() error {
	return val.ValidateStruct(
		u, //передаем нужную структуру
		//проверка поля емэйл - не пустое - соотвествует емейлу
		val.Field(&u.Email, val.Required, is.Email),
		//прописываем правило в зависимости от значения поля password, это правило возаращает фн requiredIf(bool)
		val.Field(&u.Password, val.By(requiredIf(u.EncryptedPassword == "")), val.Length(6, 100)), // проверка поля пароль - не пустое - больше 6 символов и  меньше 100
	)
}

// BeforeCreate ... зашифровываем пароль и его хэш записываем в новое поле
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encriptString(u.Password) // находим хэш пароля
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func (u *User) Sanitaze() {
	u.Password = "" // удаляем пароль
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encriptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost) // генерируем хэш из пароля
	if err != nil {
		return "", err
	}
	return string(b), nil
}
