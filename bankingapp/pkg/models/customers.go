package models

import "time"

type Customers struct {
	tableName     struct{}  `sql:"cutomer"`
	CustomerId    int       `sql:"customer_id, type:bigserial PRIMARY KEY"`
	FirstName     string    `sql:"first_name, type:text NOT NULL"`
	LastName      string    `sql:"last_name, type:text NOT NULL"`
	PanNumber     string    `sql:"pan_number,type:varchar(10) NOT NULL"`
	AadharNumber  int64     `sql:"aadhar_number, type:numeric(12) NOT NULL"`
	DOB           string    `sql:"dob, type:date NOT NULL"`
	Email         string    `sql:"email, type:text NOT NULL"`
	ContactNumber int64     `sql:"contact_number,type:numeric(10) NOT NULL"`
	Addr          string    `sql:"addr, type:text"`
	Gender        string    `sql:"gender, type:text"`
	Occupation    string    `sql:"occupation, type:text"`
	CreatedAt     time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
	FkBankId      int       `sql:"fk_bank_id, type:bigint REFERENCES banks(bank_id) "`
}
