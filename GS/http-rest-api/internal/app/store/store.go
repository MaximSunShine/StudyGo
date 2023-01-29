package store

// интерфейс для вызова у объектов нужных методов при передачи их в качестве параметров
type Store interface {
	User() UserRepository
}
