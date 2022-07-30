package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustPersonal struct {
	BankId        int    `sql:"bank_id"`
	Code          string `sql:"code"`
	Name          string `sql:"name"`
	Addr          string `sql:"addr"`
	Id            int    `sql:"id"`
	FirstName     string
	LastName      string
	PanNumber     string
	AadharNumber  int64
	DOB           string
	Email         string
	ContactNumber int64
	CustAddr      string `sql:"cust_addr"`
	Gender        string
	Occupation    string
}

func CustomerDetails(db *pg.DB, custId int) (mod CustPersonal) {

	details := new(models.Customers)
	fmt.Println(custId)
	// db.Model(details).Where("customers.customer_id = ?0", custId).Select()
	// fmt.Println(details)
	err := db.Model(details).ColumnExpr("b.id AS bank_id, b.code AS code, b.name AS name, b.addr AS addr, customers.id AS id, first_name, last_name, pan_number, aadhar_number, dob, email, contact_number, customers.addr AS cust_addr, gender, occupation").Join("JOIN bank_customer_map m ON m.fk_customer_id = customers.id").Join("JOIN banks b ON fk_bank_id = b.id").Where("customers.id = ?0", custId).Select(&mod)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// db.Model(details).ColumnExpr("b.bank_id").Join("JOIN banks b ON fk_bank_id = b.bank_id").Where("customer_id = ?0", 1).Select(&mod)
	fmt.Println(mod)
	return
}
