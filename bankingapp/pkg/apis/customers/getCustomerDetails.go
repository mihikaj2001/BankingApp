package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func GetCustomerDetails(db *pg.DB, r *gin.Engine, custId int) {

	valobj := dbpkg.CustomerDetails(db, custId)
	r.GET("/customer/details", func(c *gin.Context) {
		c.JSON(http.StatusOK, &valobj)
	})
	// fmt.Println(reflect.TypeOf(r))

}
