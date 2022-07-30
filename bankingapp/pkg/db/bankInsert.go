package db

import (
	"bank_app/pkg/models"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v10"
)

func BankInsert(db *pg.DB, newBank *models.Banks) (err error) {

	newBank.CreatedAt = time.Now()
	_, err = db.Model(newBank).Insert()
	fmt.Println(newBank)
	if err != nil {
		fmt.Print(err)
	}

	return

}
