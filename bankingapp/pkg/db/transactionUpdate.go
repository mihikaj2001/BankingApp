package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type TransactData struct {
	ACCOUNT_ID int `json:"account_id"`
	AMOUNT     int `json:"amount"`
}

type TransferData struct {
	ACCOUNT_ID     int `json:"account_id"`
	AMOUNT         int `json:"amount"`
	BENEFICIARY_ID int `json:"beneficiary_id"`
}

// func TransactionWithdraw(db *pg.DB, newTransaction *TransactData) models.Transactions {
// 	var account1 models.Accounts
// 	var trans1 models.Transactions
// 	trans1.FkAccountId = newTransaction.ACCOUNT_ID
// 	err := db.Model(&account1).ColumnExpr("current_balance, fk_ifsc_code, account_number").Where("account_id = ?0", newTransaction.ACCOUNT_ID).Select()
// 	if err != nil {
// 		fmt.Print(err)
// 		panic(err)
// 	}
// 	// fmt.Println(newTransaction)
// 	fmt.Println(account1)
// 	trans1.CreditedAmount = 0
// 	trans1.DebitedAmount = newTransaction.AMOUNT
// 	trans1.RunningBalance = account1.CurrentBalance - newTransaction.AMOUNT
// 	if trans1.RunningBalance < 0 {
// 		fmt.Println("Balance too low for transaction")
// 		return trans1
// 	}
// 	trans1.OtherPartyAccountNumber = int(account1.AccountNumber)
// 	trans1.OtherPartyBankName = "same_bank"
// 	trans1.OtherPartyBranchName = "same_branch"
// 	trans1.OtherPartyIfsc = account1.FkIfscCode
// 	// trans1.
// 	fmt.Println(trans1)

// 	_, err = db.Model(&trans1).Insert()
// 	if err != nil {
// 		fmt.Print(err.Error())
// 		panic(err)
// 	}
// 	return trans1

// }

func TransactionUpdate(db *pg.DB, newTransaction *TransactData, checker int) models.Transactions {
	var account1 models.Accounts
	var trans1 models.Transactions
	// var account2 models.Accounts
	// var trans2 models.Transactions
	trans1.FkAccountId = newTransaction.ACCOUNT_ID

	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("Error while opening tx for transaction, reason : %v\n", txErr.Error())
	}
	err := tx.Model(&account1).ColumnExpr("current_balance, fk_ifsc_code, account_number").Where("account_id = ?0", newTransaction.ACCOUNT_ID).Select()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	// fmt.Println(newTransaction)
	fmt.Println(account1)
	// trans1.CreditedAmount = 0
	// trans1.DebitedAmount = newTransaction.AMOUNT
	switch checker {
	case 1:
		trans1.RunningBalance = account1.CurrentBalance - newTransaction.AMOUNT
		trans1.CreditedAmount = 0
		trans1.DebitedAmount = newTransaction.AMOUNT
	case 2:
		trans1.RunningBalance = account1.CurrentBalance + newTransaction.AMOUNT
		trans1.CreditedAmount = newTransaction.AMOUNT
		trans1.DebitedAmount = 0
	}
	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		// panic(trans1)
		return trans1
	}
	trans1.OtherPartyAccountNumber = int(account1.AccountNumber)
	trans1.OtherPartyBankName = "same_bank"
	trans1.OtherPartyBranchName = "same_branch"
	trans1.OtherPartyIfsc = account1.FkIfscCode
	// trans1.
	fmt.Println(trans1)

	_, err = tx.Model(&trans1).Insert()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account1).Set("current_balance = ?", trans1.RunningBalance).Where("account_id = ?", newTransaction.ACCOUNT_ID).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
	tx.Commit()
	return trans1

}

func TransferUpdate(db *pg.DB, newTransaction *TransactData, newBeneficiary *BenefData) models.Transactions {
	var account1 models.Accounts
	var trans1 models.Transactions
	trans1.FkAccountId = newTransaction.ACCOUNT_ID
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("Error while opening tx for transaction, reason : %v\n", txErr.Error())
	}
	err := tx.Model(&account1).ColumnExpr("current_balance, fk_ifsc_code, account_number").Where("account_id = ?0", newTransaction.ACCOUNT_ID).Select()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	// fmt.Println(newTransaction)
	fmt.Println(account1)
	// trans1.CreditedAmount = 0
	// trans1.DebitedAmount = newTransaction.AMOUNT

	trans1.RunningBalance = account1.CurrentBalance - newTransaction.AMOUNT
	trans1.CreditedAmount = 0
	trans1.DebitedAmount = newTransaction.AMOUNT
	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		// panic(trans1)
		return trans1
	}
	trans1.OtherPartyAccountNumber = int(account1.AccountNumber)
	trans1.OtherPartyBankName = "same_bank"
	trans1.OtherPartyBranchName = "same_branch"
	trans1.OtherPartyIfsc = account1.FkIfscCode
	// trans1.
	fmt.Println(trans1)

	_, err = tx.Model(&trans1).Insert()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account1).Set("current_balance = ?", trans1.RunningBalance).Where("account_id = ?", newTransaction.ACCOUNT_ID).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	tx.Commit()
	return trans1

}
