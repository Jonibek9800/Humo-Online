package models

import (
	"ATMsExample/db"
	"database/sql"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
	Admin    bool   `json:"admin"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Remove   bool   `json:"remove"`
}

func (receiver User) AddNewUser(Db *sql.DB, user User) (err error) {
	_, err = Db.Exec(db.AddUser, user.Name, user.Surname, user.Age, user.Gender, user.Admin, user.Login, user.Password)
	if err != nil {
		return err
	}
	return
}
