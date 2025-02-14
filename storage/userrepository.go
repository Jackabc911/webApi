package storage

import (
	"fmt"
	"log"

	"github.com/Jackabc911/webApi/internal/app/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create user in database
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, hashedpassword, secretnumber) VALUES ($1, $2, $3) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(
		query,
		u.Login,
		u.HashedPassword,
		u.SecretNumber,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// Find by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

// Select All
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.HashedPassword, &u.SecretNumber)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
