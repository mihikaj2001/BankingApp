package models

import "time"

type Customers struct {
	tableName     struct{}  ` sql:"customers"`
	Id            int       `json:"id,omitempty" sql:"id, type:bigserial PRIMARY KEY"`
	FirstName     string    `json:"first_name,omitempty" sql:"first_name, type:text NOT NULL"`
	LastName      string    `json:"last_name,omitempty" sql:"last_name, type:text NOT NULL"`
	PanNumber     string    `json:"pan_number,omitempty" sql:"pan_number,type:varchar(10) NOT NULL"`
	AadharNumber  int64     `json:"aadhar_number,omitempty" sql:"aadhar_number, type:numeric(12) NOT NULL"`
	DOB           string    `json:"dob,omitempty" sql:"dob, type:date NOT NULL"`
	Email         string    `json:"email,omitempty" sql:"email, type:text NOT NULL"`
	ContactNumber int64     `json:"contact_number,omitempty" sql:"contact_number,type:numeric(10) NOT NULL"`
	Addr          string    `json:"addr,omitempty" sql:"addr, type:text"`
	Gender        string    `json:"gender,omitempty" sql:"gender, type:gender_enum"`
	Occupation    string    `json:"occupation,omitempty" sql:"occupation, type:text"`
	CreatedAt     time.Time `json:"created_at,omitempty" sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
