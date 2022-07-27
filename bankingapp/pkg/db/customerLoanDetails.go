package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type CustLoans struct {
	LoanId       int
	FkCustomerId int
	LoanAmount   int
	LoanTerm     int
	LoanInterest int
}

func CustomerLoanDetails(db *pg.DB, custId int) []CustLoans {
	var inputs models.Loans
	var results []CustLoans
	fmt.Println(custId)
	err := db.Model(&inputs).ColumnExpr("loan_id, fk_customer_id, loan_amount, loan_term, loan_interest").Where("fk_customer_id = ?0", custId).Select(&results)
	// err := db.Model(&inputs).Column("loan_id").Where("fk_customer_id = ?0", custId).Select(&results)
	if err != nil {
		panic(err)
	}

	// for i := range results {
	// 	fmt.Println(results[i])

	// }

	return results
	// return 5

}
