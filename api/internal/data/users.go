package data

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (m *UserModel) GetUserByEmail(email string) (*User, error) {
	stmt := "SELECT id, email, password FROM users WHERE email = $1"
	u := &User{}
	err := m.DB.QueryRow(context.Background(), stmt, email).Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		// if errors.Is(pgx.ErrNoRows, err) {
		// 	return nil, ErrRecordNotFound
		// }

		return nil, err
	}
	return u, nil
}

func (m *UserModel) GetUserById(id int64) (*User, error) {
	stmt := "SELECT id, email, password FROM users WHERE id = $1"
	u := &User{}
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (m *UserModel) Insert(user *User) (*User, error) {
	// hash the password before inserting into db
	hashedPwd, err := GenerateHash(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPwd

	stmt := "INSERT INTO users (email, password, firstName, lastName) VALUES ($1, $2, $3, $4) RETURNING id"

	err = m.DB.QueryRow(context.Background(), stmt, user.Email, user.Password, user.FirstName, user.LastName).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GenerateHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
