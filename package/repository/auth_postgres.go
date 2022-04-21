package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/rest-api/internal/user"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(u user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, name, surname, email, password_hash) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := ap.db.QueryRow(query, u.Username, u.Name, u.Surname, u.Email, u.Password)
	if err := row.Scan(&id); err != nil {
		logrus.Fatalf("Cannot scan id, due to error: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (user.User, error) {
	var u user.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := ap.db.Get(&u, query, username, password)

	return u, err
}
