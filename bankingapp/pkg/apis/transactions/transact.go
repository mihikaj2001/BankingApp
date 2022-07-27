package transactions

import (
	dbpkg "bank_app/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Transact(db *pg.DB, r *gin.Engine, checker int) {
	r.POST("/transaction/transact", func(c *gin.Context) {
		newTransaction := &dbpkg.TransactData{}
		// login.USER = "mihika"
		c.ShouldBindJSON(&newTransaction)
		c.JSON(200, dbpkg.TransactionUpdate(db, newTransaction, checker)) // Your custom response here
	})
}
