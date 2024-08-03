package repository

import (
	"database/sql"
	"fmt"
	"quiz-3/model"
)

type UserRepository interface {
	GetUserByUsernameAndPassword(username string, password string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUsernameAndPassword(username string, password string) (*model.User, error) {
	var user model.User
	sql := "SELECT * FROM users WHERE username = $1 AND password = $2"
	err := r.db.QueryRow(sql, username, password).Scan(&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.ModifiedAt,
		&user.ModifiedBy)
	fmt.Println("Cek get user:", err)
	// err := r.db.QueryRow(sql, username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
