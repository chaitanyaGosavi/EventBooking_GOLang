package models

import (
	"errors"
	"eventsManagement/db"
	"eventsManagement/utils"
	"fmt"
)

type User struct {
	Id       int64
	Name     string `binding : "required`
	Email    string `binding : "required`
	Password string `binding : "required`
}

func (u User) CreateNewUser() error {
	createUserQuery := `
	INSERT INTO users (Name, Email, Password)
	VALUES (?, ?, ?);
	`

	statement, err := db.DB.Prepare(createUserQuery)
	if err != nil {
		return err
	}
	fmt.Println("Statement Prepared")
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	fmt.Println("Pasword hashed")
	defer statement.Close()
	result, err := statement.Exec(u.Name, u.Email, hashedPassword)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Statement Executed")
	id, err := result.LastInsertId()

	u.Id = id
	fmt.Println("Statement Executed2")
	return err

}

func (u *User) ValidateCredentials() error {

	getPasswordQuery := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(getPasswordQuery, u.Email)

	var dbSavedPassword string
	err := row.Scan(&u.Id, &dbSavedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckHashWithPassword(u.Password, dbSavedPassword)

	if !isPasswordValid {
		return errors.New("Invalid Credentials")
	}

	return nil
}
