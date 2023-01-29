package teststore

import (
	"github.com/GS/http-rest-api/internal/app/model"
	"github.com/GS/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

// создание подобия БД (хэш/мапу), чтобы не нагружать БД
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		// создаем мапу где ключ будет емэйл, а строка таблицы будет значением в виде структуры model.User
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
