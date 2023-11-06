package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountId      int    `json:"account_id" sql:"account_id, type:int PRIMARY KEY"`
	FkCustomerId   int    `json:"fk_customer_id" sql:"fk_customer_id, type:int REFERENCES customer(customer_id) "`
	AccountNumber  int64  `json:"account_number" sql:"account_number, type:numeric UNIQUE NOT NULL"`
	IsActive       bool   `json:"is_active" sql:"is_active, type:boolean"`
	AccountType    string `json:"account_type" sql:"account_type, type:text NOT NULL"`
	CurrentBalance int    `json:"current_balance" sql:"current_balance, type:numeric(10,2) NOT NULL"`
}
