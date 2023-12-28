package data

import (
	"context"
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
		// if errors.Is(pgx.ErrNoRows, err) {
		// 	return nil, ErrRecordNotFound
		// }

		return nil, err
	}
	return u, nil
}
