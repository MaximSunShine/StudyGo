package teststore

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// создание нового пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	} //проверка на верность заполненных полей в соотвествии с правилами

	if err := u.BeforeCreate(); err != nil {
		return err
	} //зашифровываем пароль и его хэш записываем в новое поле

	r.users[u.Email] = u // добавляем в мапу новый? ключ и его значение
	u.ID = len(r.users)  // сомнительно

	return nil
}

// поиск пользователя по емэйлу
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email] // если в пмапе по ключу есть такой пользователь, то возвращаем его
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
