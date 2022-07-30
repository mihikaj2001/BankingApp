package models

import "time"

type Banks struct {
	tableName struct{}  `sql:"banks"`
	Id        int       `json:",omitempty" sql:"id, type:bigserial PRIMARY KEY"`
	Code      string    `json:",omitempty" sql:"code, type: text UNIQUE"`
	Name      string    `json:",omitempty" sql:"name, type:text"`
	Addr      string    `json:",omitempty" sql:"addr, type:text"`
	CreatedAt time.Time `json:",omitempty" sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:",omitempty" sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
