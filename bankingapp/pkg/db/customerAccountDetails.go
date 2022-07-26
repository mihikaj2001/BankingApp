package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustAccounts struct {
	Id             int
	FkCustomerId   int
	FkBranchId     int
	AccountNumber  int64
	IsActive       bool
	AccountType    string
	CurrentBalance int
}

func CustomerAccountDetails(db *pg.DB, custId int) []CustAccounts {
	var inputs models.Accounts
	var results []CustAccounts
	fmt.Println(custId)
	err := db.Model(&inputs).ColumnExpr("id, fk_customer_id, fk_branch_id, account_number, is_active, account_type, current_balance").Where("fk_customer_id = ?0", custId).Select(&results)
	// err := db.Model(&results).ColumnExpr("account_id").Where("fk_customer_id = ?0", custId).Select()
	if err != nil {
		panic(err)
	}

	for i := range results {
		fmt.Println(results[i])

	}

	return results
	// return 5

}
