package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerId       int    `sql:"customer_id, type:int PRIMARY KEY"`
	Name             string `json:"name" sql:"name, type:text NOT NULL"`
	CPF_CNPJ         string `json:"CPF OR CNPJ" sql:"CPF OR CNPJ, type: text NOT NULL"`
	PhoneNumber      string `json:"phonenumber" sql:"phonenumber, type: text NOT NULL"`
	Email            string `json:"email" sql:"email, type:text NOT NULL"`
	FkCurrentBalance int    `json:"fk_current_balance`
	FkBankId         int    `json:"fk_bank_id" sql:"fk_bank_id, type:int REFERENCES bank(bank_id) "`
	Check            int    `json:"transaction_type"`
}
