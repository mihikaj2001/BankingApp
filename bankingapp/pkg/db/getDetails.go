package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

func GetDetails(db *pg.DB, custId int) models.Customers {
	details := new(models.Customers)
	fmt.Println(details)
	db.Model(details).Where("customer_id = ?0", custId).Select()
	fmt.Println(details)
	return *details
}
