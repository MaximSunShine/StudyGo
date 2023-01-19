package store

import "github.com/GS/http-rest-api/internal/app/model"

// UserRepository создаем структуру UserRepository которая будет полем структуры Store(основаная структура БД)
// к объкту UserRepository мы будем обращаться через метод Store.User.{...}, так как его имя с маленькой буквы и мы при
// объявлении новой структуры Store не будем его видеть через . !
// У каждой таблици будет своя структура и свой доп объект выполнения разных команд к таблице посредством методов объекта
// а обращаться мы к объекту будем через метот главной структуры БД именованной именем конкретной таблицы
type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1,$2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
