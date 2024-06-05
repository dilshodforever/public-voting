package postgres

import (
	"database/sql"
	"fmt"
	"root/config"
	u "root/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	Publics u.Public
	Partys  u.Party
}

func NewPostgresStorage() (u.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, err

}

func (s *Storage) Public() u.Public {
	if s.Publics == nil {
		s.Publics = &PublicStorage{s.Db}
	}
	return s.Publics
}

func (s *Storage) Party() u.Party {
	if s.Partys == nil {
		s.Partys = &PartyStorage{s.Db}
	}
	return s.Partys
}




