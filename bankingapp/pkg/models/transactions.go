package models

import "time"

type Transactions struct {
	tableName           struct{}  `sql:"transactions"`
	Id                  int       `sql:"id,type:bigserial PRIMARY KEY"`
	FkAccountId         int       `sql:"fk_account_id, type:bigint REFERENCES accounts(id) ON DELETE CASCADE"`
	Amount              int       `sql:"amount, type:int"`
	TransactionType     string    `sql:"transaction_type, type:transaction_type_enum"`
	RunningBalance      int       `sql:"running_balance, type:int"`
	OtherPartyAccountId int       `sql:"other_party_account_id, type:numeric(10) NOT NULL"`
	TransactionAt       time.Time `sql:"transaction_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
