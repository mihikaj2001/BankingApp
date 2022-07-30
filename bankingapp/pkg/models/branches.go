package models

import "time"

type Branches struct {
	tableName struct{}  `sql:"branches"`
	Id        int       `sql:"id, type:bigserial PRIMARY KEY"`
	IfscCode  string    `sql:"ifsc_code, type:varchar(11) UNIQUE NOT NULL"`
	FkBankId  int       `sql:"fk_bank_id, type:bigint REFERENECES banks(id)" `
	Name      string    `sql:"name, type:text"`
	Addr      string    `sql:"addr, type:text"`
	CreatedAt time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
