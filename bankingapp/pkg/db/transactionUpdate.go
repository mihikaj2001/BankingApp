package db

import (
	"bank_app/pkg/models"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v10"
)

type TransactData struct {
	ACCOUNT_ID int `json:"id"`
	AMOUNT     int `json:"amount"`
}

type TransferData struct {
	ACCOUNT_ID     int `json:"account_id"`
	AMOUNT         int `json:"amount"`
	BENEFICIARY_ID int `json:"beneficiary_id"`
}

// FOR TRANSACTION
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
	err := tx.Model(&account1).ColumnExpr("current_balance, fk_branch_id, account_number").Where("id = ?0", newTransaction.ACCOUNT_ID).Select()
	if err != nil {
		fmt.Print(err)
		_ = tx.Rollback()
	}
	// fmt.Println(newTransaction)
	fmt.Println(account1)
	// trans1.CreditedAmount = 0
	// trans1.DebitedAmount = newTransaction.AMOUNT
	switch checker {
	case 1:
		trans1.RunningBalance = account1.CurrentBalance - newTransaction.AMOUNT
		trans1.Amount = newTransaction.AMOUNT
		trans1.TransactionType = "Withdrawal"
	case 2:
		trans1.RunningBalance = account1.CurrentBalance + newTransaction.AMOUNT
		trans1.Amount = newTransaction.AMOUNT
		trans1.TransactionType = "Deposit"
	}
	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		// panic(trans1)
		return trans1
	}
	trans1.OtherPartyAccountId = newTransaction.ACCOUNT_ID
	fmt.Println(trans1)

	_, err = tx.Model(&trans1).Insert()
	if err != nil {
		fmt.Print(err.Error())
		_ = tx.Rollback()
	}

	_, err = tx.Model(&account1).Set("current_balance = ?0, updated_at = ?1", trans1.RunningBalance, time.Now()).Where("id = ?", newTransaction.ACCOUNT_ID).Update()
	if err != nil {
		fmt.Print(err.Error())
		_ = tx.Rollback()
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}
	return trans1

}

// FOR TRANSFERS
func TransferUpdate(db *pg.DB, newTransfer *TransferData) models.Transactions {
	// INITIALIZATION
	var account1 models.Accounts
	var trans1 models.Transactions
	var account2 models.Accounts
	var trans2 models.Transactions

	// SETTING TRANSACTION ACCOUNT_ID
	trans1.FkAccountId = newTransfer.ACCOUNT_ID
	trans2.FkAccountId = newTransfer.BENEFICIARY_ID

	// BEGIN TRANSACTION
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("Error while opening tx for transaction, reason : %v\n", txErr.Error())
	}
	// GET ACCOUNT1 DATA
	err := tx.Model(&account1).ColumnExpr("current_balance, fk_branch_id, account_number").Where("id = ?0", newTransfer.ACCOUNT_ID).Select()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	// GET ACCOUNT2 DATA
	err = tx.Model(&account2).ColumnExpr("current_balance, fk_branch_id, account_number").Where("id = ?0", newTransfer.BENEFICIARY_ID).Select()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	fmt.Println(account1)
	fmt.Println(account2)

	trans1.RunningBalance = account1.CurrentBalance - newTransfer.AMOUNT
	trans2.RunningBalance = account1.CurrentBalance + newTransfer.AMOUNT

	trans1.Amount = newTransfer.AMOUNT
	trans2.Amount = newTransfer.AMOUNT

	trans1.TransactionType = "Withdrawal"
	trans2.TransactionType = "Deposit"

	trans1.OtherPartyAccountId = newTransfer.BENEFICIARY_ID
	trans2.OtherPartyAccountId = newTransfer.ACCOUNT_ID

	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		// panic(trans1)
		return trans1
	}

	fmt.Println(trans1)
	fmt.Println(trans2)

	// INSERTING VALUES INTO TRANSACTION TABLE
	_, err = tx.Model(&trans1).Insert()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&trans2).Insert()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account1).Set("current_balance = ?0, updated_at = ?1", trans1.RunningBalance, time.Now()).Where("id = ?", newTransfer.ACCOUNT_ID).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account2).Set("current_balance = ?0, updated_at = ?1", trans2.RunningBalance, time.Now()).Where("id = ?", newTransfer.BENEFICIARY_ID).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	tx.Commit()
	return trans1

}
