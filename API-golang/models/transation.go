package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionId           int `sql:"transaction_id,type:int PRIMARY KEY"`
	FkAccountId             int `sql:"fk_account_id, type:int REFERENCES account(account_id)"`
	CreditedAmount          int `sql:"credited_amount, type:numeric(10,2)"`
	DebitedAmount           int `sql:"debited_amount, type:numeric(10,2)"`
	RunningBalance          int `sql:"running_balance, type:numeric(10,2)"`
	OtherPartyAccountNumber int `sql:"other_party_account_number, type:int "`
	Check                   int `json:"transaction_type"`
}
