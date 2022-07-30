package models

import "time"

type Loans struct {
	tableName       struct{}  `sql:"loans"`
	Id              int       `json:",omitempty" sql:"id, type:bigserial PRIMARY KEY"`
	FkCustomerId    int       `sql:"fk_customer_id, type:bigint REFERENCES customers(id) ON DELETE CASCADE"`
	FkBranchId      int       `json:",omitempty" sql:"fk_branch_id, type:bigint REFERENCES branches(id) ON DELETE CASCADE"`
	Amount          int       `sql:"amount, type:bigint"`
	Term            int       `sql:"term, type:int"`
	InterestPercent int       `sql:"interest_percent, type:int"`
	TotalInterest   int       `sql:"total_interest, type:int"`
	Installments    int       `sql:"installments, type: int"`
	MonthlyAmount   int       `sql:"monthly_amount, type:int"`
	MonthlyInterest int       `sql:"monthly_interest, type:int"`
	CreatedAt       time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
