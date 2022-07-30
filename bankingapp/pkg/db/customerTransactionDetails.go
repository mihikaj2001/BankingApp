package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustTransaction struct {
	Id              int
	FkAccountId     int
	Amount          int
	TransactionType string
	RunningBalance  int
}

func CustomerTransactionDetails(db *pg.DB, accountId int) []CustTransaction {
	var inputs models.Transactions
	var results []CustTransaction
	fmt.Println(accountId)
	err := db.Model(&inputs).ColumnExpr("id, fk_account_id, amount, transaction_type, running_balance").Where("fk_account_id = ?0", accountId).Select(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	return results
}
