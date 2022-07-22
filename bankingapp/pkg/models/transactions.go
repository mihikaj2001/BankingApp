package models

import "time"

type Transactions struct {
	tableName               struct{}  `sql:"transactions"`
	TransactionId           int       `sql:"transaction_id,type:bigserial PRIMARY KEY"`
	FkAccountId             int       `sql:"fk_account_id, type:bigint REFERENCES accounts(account_id)"`
	CreditedAmount          int       `sql:"credited_amount, type:int"`
	DebitedAmount           int       `sql:"debited_amount, type:int"`
	RunningBalance          int       `sql:"running_balance, type:int"`
	OtherPartyIfsc          string    `sql:"other_party_ifsc, type:varchar(11) NOT NULL"`
	OtherPartyAccountNumber int       `sql:"other_party_account_number, type:numeric(10) NOT NULL"`
	OtherPartyBankName      string    `sql:"other_party_bank_name, type:text NOT NULL"`
	OtherPartyBranchName    string    `sql:"other_party_branch_name, type:text NOT NULL"`
	TransactionAt           time.Time `sql:"transaction_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
