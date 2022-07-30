package db

import (
	"bank_app/pkg/models"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v10"
)

func AccountInsert(db *pg.DB, newAccount *models.Accounts) (err error) {

	newAccount.CreatedAt = time.Now()
	_, err = db.Model(newAccount).Insert()
	fmt.Println(newAccount)
	if err != nil {
		fmt.Print(err)
	}

	return

}
