package customers

import (
	dbpkg "bank_app/pkg/db"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func GetCustomerDetails(db *pg.DB, r *gin.Engine) {
	valtest := dbpkg.GetDetails(db, 1)
	r.GET("/customer/details", func(c *gin.Context) {
		c.JSON(http.StatusOK, &valtest)
	})
	fmt.Println(reflect.TypeOf(r))

}
