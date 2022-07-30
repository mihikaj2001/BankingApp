package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func GetCustomerTransactions(db *pg.DB, r *gin.Engine, accountId int) {
	valobj := dbpkg.CustomerTransactionDetails(db, accountId)
	r.GET("/customer/transactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, &valobj)
	})

}
