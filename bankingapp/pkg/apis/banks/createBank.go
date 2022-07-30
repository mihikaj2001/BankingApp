package banks

import (
	dbpkg "bank_app/pkg/db"
	"bank_app/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func CreateBank(db *pg.DB, r *gin.Engine) {
	r.POST("/banks/insert", func(c *gin.Context) {
		newBank := &models.Banks{}
		c.ShouldBindJSON(&newBank)
		c.JSON(200, dbpkg.BankInsert(db, newBank))
	})
}
