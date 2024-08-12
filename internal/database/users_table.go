package database

import (
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func (dm *DBManager) CreateUser(user User) (int, error) {
	var id int
	err := dm.db.QueryRow("insert into users (email,username,password) values ($1,$2,$3) returning id", user.Email, user.Name, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dm *DBManager) VerifyPassword(email, password string) error {
	var hashedPassword []byte
	err := dm.db.QueryRow("select password from users where email = $1", email).Scan(&hashedPassword)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
