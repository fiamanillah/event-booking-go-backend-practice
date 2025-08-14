package models

import (
	"errors"
	"gin/db"
	"gin/utils"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)
	u.Password = hashedPassword
	return err
}

func (u *User) ValidateCredentials(email, password string) error {
	query := "SELECT   email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	var hashedPassword string
	err := row.Scan(&u.Email, &hashedPassword)
	if err != nil {
		return errors.New("invalid email or password")
	}

	passwordIsValid := utils.CheckPassword(password, hashedPassword)
	println(passwordIsValid)
	if !passwordIsValid {
		return errors.New("invalid email or password")
	}
	return nil
}
