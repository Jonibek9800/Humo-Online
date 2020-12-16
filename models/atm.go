package models

import (
	"ATMsExample/db"
	"database/sql"
)

type ATM struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	Status  bool   `json:"status"`
}

func (receiver ATM) AddNewATM(Db *sql.DB, atm ATM) (err error) {
	_, err = Db.Exec(db.AddATM, atm.Address)
	if err != nil {
		return err
	}
	return
}
func AddNewATM(Db *sql.DB, address string) (err error) {
	_, err = Db.Exec(db.AddATM, address)
	if err != nil {
		return err
	}
	return
}
