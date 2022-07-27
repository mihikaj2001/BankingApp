package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustTransaction struct {
	TransactionId           int
	FkAccountId             int
	CreditedAmount          int
	DebitedAmount           int
	RunningBalance          int
	OtherPartyIfsc          string
	OtherPartyAccountNumber int
	OtherPartyBankName      string
	OtherPartyBranchName    string
}

func CustomerTransactionDetails(db *pg.DB, accountId int) []CustTransaction {
	var inputs models.Transactions
	var results []CustTransaction
	fmt.Println(accountId)
	err := db.Model(&inputs).ColumnExpr("transaction_id, fk_account_id, credited_amount, debited_amount, running_balance , other_party_ifsc, other_party_account_number, other_party_bank_name, other_party_branch_name").Where("fk_account_id = ?0", accountId).Select(&results)
	if err != nil {
		panic(err)
	}
	return results
}
