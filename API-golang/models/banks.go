package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	BankId      int    `sql:"bank_id, type:int PRIMARY KEY"`
	BankName    string `sql:"bank_name, type:text"`
	BankCode    int    `sql:"bank_code, type: int UNIQUE"`
	BankAgency  string `sql:"bank_agency, type:text"`
	BankAccount string `sql:"bank_account, type:text"`
}
