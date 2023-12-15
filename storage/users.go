package storage

import (
	"github.com/dgrijalva/jwt-go"
	"hackaton/types"
)

func (st *PStorage) SearchUser(user *types.User, token *jwt.StandardClaims) error {
	err := st.Db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", token.Issuer).
		Scan(&user.Id, &user.Name, &user.Email)
	return err
}

func (st *PStorage) AddUser(user *types.User) error {
	_, err := st.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Email, user.Password)
	if err != nil {
		return err
	}

	return err
}
