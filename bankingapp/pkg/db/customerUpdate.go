package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustUpdate struct {
	ID      int64 `json:"customer_id"`
	CONTACT int64 `json:"contact_number"`
}

func CustomerUpdate(db *pg.DB, newUpdate *CustUpdate) models.Customers {
	var customer models.Customers
	fmt.Println(newUpdate.CONTACT)
	fmt.Println(newUpdate.ID)
	_, err := db.Model(&customer).Set("contact_number = ?", newUpdate.CONTACT).Where("customer_id = ?", newUpdate.ID).Update()
	fmt.Printf("err %#v", err)
	if err != nil {
		panic(err)
	}
	return customer
}
