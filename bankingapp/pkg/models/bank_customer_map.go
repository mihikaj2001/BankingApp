package models

type BankCustomerMap struct {
	tableName    struct{} `sql:"bank_customer_map"`
	Id           int      `json:",omitempty" sql:"id, type:bigserial PRIMARY KEY"`
	FkBankId     int      `sql:"fk_bank_id, type:bigint REFERENECES banks(id) ON DELETE CASCADE" `
	FkCustomerId int      `json:"fk_customer_id" sql:"fk_customer_id, type:bigint REFERENCES customers(id) ON DELETE CASCADE"`
}
