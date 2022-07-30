package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func GetCustomerAccountDetails(db *pg.DB, r *gin.Engine, custId int) {
	valobj := dbpkg.CustomerAccountDetails(db, custId)
	r.GET("/customer/accounts", func(c *gin.Context) {
		for i := range valobj {
			c.JSON(http.StatusOK, valobj[i])
		}
		// c.JSON(http.StatusOK, valobj)

	})
}
