package accounts

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

func CreateAccount(db *pg.DB, r *gin.Engine) {
	// valobj := dbpkg.InsertCustomerDetails(db, accountId)
	r.POST("/account/insert", func(c *gin.Context) {
		newAccount := &models.Accounts{}
		// login.USER = "mihika"
		c.ShouldBindJSON(&newAccount)
		c.JSON(200, dbpkg.AccountInsert(db, newAccount)) // Your custom response here
	})
}
