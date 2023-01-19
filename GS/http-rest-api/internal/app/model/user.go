package model

type User struct { // создаем аналогию таблицы базы данных
	ID                int
	Email             string
	EncryptedPassword string
}
