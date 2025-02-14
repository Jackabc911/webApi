package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Instance of store
type Storage struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// Constructor for storage
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open storage method
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	//Проверим, что все ок. Реально соединение тут не создается. Соединение только при первом вызове
	//db.Ping() // Пустой SELECT *
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Connection to db successfully")
	return nil
}

// Public for UserRepo
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
