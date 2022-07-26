package models

import "time"

type Accounts struct {
	tableName      struct{}  `sql:"accounts"`
	AccountId      int       `sql:"account_id, type:bigserial PRIMARY KEY"`
	FkCustomerId   int       `sql:"fk_customer_id, type:bigint REFERENCES customers(customer_id) ON DELETE CASCADE"`
	AccountNumber  int64     `sql:"account_number, type:numeric(10) UNIQUE NOT NULL"`
	IsActive       bool      `sql:"is_active, type:boolean"`
	AccountType    string    `sql:"account_type, type:text NOT NULL"`
	CurrentBalance int       `sql:"current_balance, type:bigint NOT NULL"`
	FkIfscCode     string    `sql:"fk_ifsc_code, type:varchar(11) REFERENCES branches(ifsc_code) NOT NULL"`
	CreatedAt      string    `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
