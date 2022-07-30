package models

import "time"

type Accounts struct {
	tableName      struct{}  `sql:"accounts"`
	Id             int       `json:"id,omitempty" sql:"id, type:bigserial PRIMARY KEY"`
	FkCustomerId   int       `json:"fk_customer_id" sql:"fk_customer_id, type:bigint REFERENCES customers(id) ON DELETE CASCADE"`
	FkBranchId     int       `json:"fk_branch_id,omitempty" sql:"fk_branch_id, type:bigint REFERENCES branches(id) ON DELETE CASCADE"`
	AccountNumber  int       `json:"account_number" sql:"account_number, type:numeric(10) UNIQUE NOT NULL"`
	IsActive       bool      `json:"is_active" sql:"is_active, type:boolean"`
	AccountType    string    `json:"account_type" sql:"account_type, type:account_type_enum NOT NULL"`
	CurrentBalance int       `json:"current_balance" sql:"current_balance, type:bigint NOT NULL"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
