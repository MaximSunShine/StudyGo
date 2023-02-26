package teststore

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// создание нового пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	} //проверка на верность заполненных полей в соотвествии с правилами

	if err := u.BeforeCreate(); err != nil {
		return err
	} //зашифровываем пароль и его хэш записываем в новое поле

	u.ID = len(r.users) + 1 // сомнительно
	r.users[u.ID] = u       // добавляем в мапу новый? ключ и его значение
	
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id] // если в пмапе по ключу есть такой пользователь, то возвращаем его
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// поиск пользователя по емэйлу
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
