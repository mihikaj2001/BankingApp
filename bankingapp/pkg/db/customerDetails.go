package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustPersonal struct {
	BankId        int    `sql:"bank_id`
	BankCode      string `sql:"bank_code"`
	BankName      string `sql:"bank_name"`
	BankAddress   string `sql:"bank_address"`
	CustomerId    int    `sql:"customer_id "`
	FirstName     string `sql:"first_name "`
	LastName      string `sql:"last_name "`
	PanNumber     string `sql:"pan_number"`
	AadharNumber  int64  `sql:"aadhar_number"`
	DOB           string `sql:"dob "`
	Email         string `sql:"email "`
	ContactNumber int64  `sql:"contact_number"`
	Addr          string `sql:"addr"`
	Gender        string `sql:"gender"`
	Occupation    string `sql:"occupation"`
}

func CustomerDetails(db *pg.DB, custId int) (mod CustPersonal) {

	details := new(models.Customers)
	fmt.Println(custId)
	// db.Model(details).Where("customers.customer_id = ?0", custId).Select()
	// fmt.Println(details)
	err := db.Model(details).ColumnExpr("b.bank_id, b.bank_code, b.bank_name, b.bank_address, customer_id, first_name, last_name, pan_number, aadhar_number, dob, email, contact_number, addr, gender, occupation").Join("JOIN banks b ON fk_bank_id = b.bank_id").Where("customer_id = ?0", custId).Select(&mod)
	if err != nil {
		panic(err)
	}
	// db.Model(details).ColumnExpr("b.bank_id").Join("JOIN banks b ON fk_bank_id = b.bank_id").Where("customer_id = ?0", 1).Select(&mod)
	fmt.Println(mod)
	return
}
