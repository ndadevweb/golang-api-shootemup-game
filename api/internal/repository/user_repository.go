package repository

import (
	"database/sql"
	"errors"
	"gowebgame/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	stmt, err := r.DB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password)
	if err != nil {
		return errors.New("username already exists or DB error")
	}

	return nil
}
