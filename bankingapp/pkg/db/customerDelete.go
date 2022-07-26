package db

import (
	"bank_app/pkg/models"

	pg "github.com/go-pg/pg/v10"
)

func CustomerDelete(db *pg.DB, custId int) models.Customers {
	var result models.Customers

	_, err := db.Model(&result).Where("customer_id = ?", custId).Delete()
	if err != nil {
		panic(err)
	}

	return result
	// return 5

}
