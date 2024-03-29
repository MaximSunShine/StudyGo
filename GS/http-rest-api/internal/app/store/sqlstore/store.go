package sqlstore

import (
	"database/sql"
	"github.com/GS/http-rest-api/internal/app/store"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Store объявление структуры Store управление всеми процессами базы данных и ее таблицами
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// NewStore Фн иннициализации структуры Store и возврат ее объекта
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

/*
// Open создание подключения к базе данных и проверка пинга(тестового запроса)
func (s *Store) Open() error {

	db, err := sql.Open("pgx", s.config.DatabaseUrl)
	fmt.Println("dbUrl -> " + s.config.DatabaseUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close() //закрытие подключения к БД
}
*/
// User создание нового объекта (структуры) userRepository если у объекта Store он еще не иннициализирован, который будет отвечать за
// работу с таблицей User_table(users)
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{ //привязываем объект UserRepository к Store как часть его поля
		store: s,
	}
	return s.userRepository
}
