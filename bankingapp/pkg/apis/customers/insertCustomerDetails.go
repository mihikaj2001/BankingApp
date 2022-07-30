package customers

import (
	dbpkg "bank_app/pkg/db"

	"bank_app/pkg/models"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

// type CustomerData struct {
// 	USER     string `json:"user" binding:"required"`
// 	PASSWORD string `json:"password" binding:"required"`
// }

func InsertCustomerDetails(db *pg.DB, r *gin.Engine) {
	// valobj := dbpkg.InsertCustomerDetails(db, accountId)
	r.POST("/customer/insert", func(c *gin.Context) {
		newUser := &models.Customers{}
		// login.USER = "mihika"
		c.ShouldBindJSON(&newUser)
		c.JSON(200, dbpkg.CustomerInsert(db, newUser)) // Your custom response here
	})
}
