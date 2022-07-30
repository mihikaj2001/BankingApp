package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func DeleteCustomerDetails(db *pg.DB, r *gin.Engine, custId int) {
	r.DELETE("/customer/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, dbpkg.CustomerDelete(db, custId))

	})
}
