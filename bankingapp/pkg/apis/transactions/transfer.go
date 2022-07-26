package transactions

import (
	dbpkg "bank_app/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Transfer(db *pg.DB, r *gin.Engine) {
	// using post not put so that it adds new data instead of just updating those with already existing similar data in the database
	r.POST("/transaction/transfer", func(c *gin.Context) {
		newTransfer := &dbpkg.TransferData{}
		// login.USER = "mihika"
		c.ShouldBindJSON(&newTransfer)
		c.JSON(200, dbpkg.TransferUpdate(db, newTransfer)) // Your custom response here
	})

}
