package db

import (
	"bank_app/pkg/models"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v10"
)

type CustLoan struct {
	ID               int `json:"id"`
	AMOUNT           int `json:"amount"`
	TERM             int `json:"term"`
	INTEREST_PERCENT int `json:"interest_percent"`
	BRANCH_ID        int `json:"branch_id"`
	ACCOUNT_NUMBER   int `json:"account_number"`
}

func LoanInsert(db *pg.DB, newLoan *CustLoan) (err error) {
	// fmt.Println(newLoan)
	loan1 := new(models.Loans)
	account1 := new(models.Accounts)

	tx, err := db.Begin()
	// loan1.Id = 2
	loan1.FkCustomerId = newLoan.ID
	account1.FkCustomerId = newLoan.ID

	loan1.FkBranchId = newLoan.BRANCH_ID
	account1.FkBranchId = newLoan.BRANCH_ID

	loan1.Amount = newLoan.AMOUNT
	account1.CurrentBalance = newLoan.AMOUNT

	account1.AccountNumber = newLoan.ACCOUNT_NUMBER
	account1.AccountType = "Loan"

	account1.IsActive = true
	account1.CreatedAt = time.Now()
	account1.UpdatedAt = time.Now()

	loan1.Term = newLoan.TERM
	loan1.InterestPercent = newLoan.INTEREST_PERCENT
	loan1.TotalInterest = (newLoan.AMOUNT * newLoan.TERM * newLoan.INTEREST_PERCENT) / 100
	loan1.Installments = 12 * newLoan.TERM
	loan1.MonthlyInterest = loan1.TotalInterest / loan1.Installments
	loan1.MonthlyAmount = (loan1.Amount / loan1.Installments) + loan1.MonthlyInterest
	loan1.CreatedAt = time.Now()
	loan1.UpdatedAt = time.Now()
	_, err = tx.Model(loan1).Returning("fk_customer_id, amount, term, interest_percent, total_interest, installments, monthly_amount, monthly_interest, created_at, updated_at").Insert()
	if err != nil {
		fmt.Print(err)
		_ = tx.Rollback()
		return err
	}
	// fmt.Println(loan1)

	_, err = tx.Model(account1).Returning("fk_customer_id, fk_branch_id,account_number, is_active, account_type, current_balance, created_at, updated_at").Insert()
	if err != nil {
		fmt.Print(err)
		_ = tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}
	return err
}
