package store

import "github.com/GS/http-rest-api/internal/app/model"

// интерфейс для вызова у объектов нужных методов при передачи их в качестве параметров
type UserRepository interface {
	Create(user *model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
