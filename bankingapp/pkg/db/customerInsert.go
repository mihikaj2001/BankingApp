package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

func CustomerInsert(db *pg.DB, cust *models.Customers) models.Customers {
	_, err := db.Model(cust).Insert()
	fmt.Println(cust)
	if err != nil {
		fmt.Print(err)
	}

	return *cust
}
