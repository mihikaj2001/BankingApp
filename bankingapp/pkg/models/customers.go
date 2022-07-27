package models

import "time"

type Customers struct {
	tableName     struct{}  ` sql:"customers"`
	CustomerId    int       `json:"customer_id" sql:"customer_id, type:bigserial PRIMARY KEY"`
	FirstName     string    `json:"first_name" sql:"first_name, type:text NOT NULL"`
	LastName      string    `json:"last_name" sql:"last_name, type:text NOT NULL"`
	PanNumber     string    `json:"pan_number" sql:"pan_number,type:varchar(10) NOT NULL"`
	AadharNumber  int64     `json:"aadhar_number" sql:"aadhar_number, type:numeric(12) NOT NULL"`
	DOB           string    `json:"dob" sql:"dob, type:date NOT NULL"`
	Email         string    `json:"email" sql:"email, type:text NOT NULL"`
	ContactNumber int64     `json:"contact_number" sql:"contact_number,type:numeric(10) NOT NULL"`
	Addr          string    `json:"addr" sql:"addr, type:text"`
	Gender        string    `json:"gender" sql:"gender, type:text"`
	Occupation    string    `json:"occupation" sql:"occupation, type:text"`
	CreatedAt     time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
	FkBankId      int       `json:"fk_bank_id" sql:"fk_bank_id, type:bigint REFERENCES banks(bank_id) "`
}
