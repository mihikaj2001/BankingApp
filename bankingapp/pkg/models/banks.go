package models

import "time"

type Banks struct {
	tableName   struct{}  `sql:"bank"`
	BankId      int       `sql:"bank_id, type:bigserial PRIMARY KEY"`
	BankCode    string    `sql:"bank_code, type: text UNIQUE"`
	BankName    string    `sql:"bank_name, type:text"`
	BankAddress string    `sql:"bank_address, type:text"`
	CreatedAt   time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
