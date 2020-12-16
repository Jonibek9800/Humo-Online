package models

import (
	"ATMsExample/db"
	"database/sql"
)

type Account struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Number   int64  `json:"number"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (receiver Account) AddNewAccount(Db *sql.DB, account Account) (err error) {
	_, err = Db.Exec(db.AddNewAccount, account.UserID, account.Name, account.Number, account.Amount, account.Currency)
	if err != nil {
		return err
	}
	return nil
}
