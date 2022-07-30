package loan

import (
	dbpkg "bank_app/pkg/db"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func InsertLoan(db *pg.DB, r *gin.Engine) {
	// valobj := dbpkg.InsertCustomerDetails(db, accountId)
	r.POST("/loan/insert", func(c *gin.Context) {
		newLoan := &dbpkg.CustLoan{}
		// login.USER = "mihika"
		c.ShouldBindJSON(&newLoan)
		c.JSON(200, dbpkg.LoanInsert(db, newLoan)) // Your custom response here
	})
}
