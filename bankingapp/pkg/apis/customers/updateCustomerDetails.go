package customers

import (
	dbpkg "bank_app/pkg/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

// func UpdateCustomerDetails(db *pg.DB, r *gin.Engine) {
// 	r.PUT("/customer/update/phone", func(c *gin.Context) {

// 		phoneUpdate := &dbpkg.PhoneUpdate{}
// 		if err := c.ShouldBindJSON(&phoneUpdate); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// fmt.Println(phoneUpdate)
// 		// fmt.Println(&phoneUpdate.CONTACT)
// 		// fmt.Println(phoneUpdate.ID)
// 		c.JSON(http.StatusOK, dbpkg.CustomerUpdate(db, phoneUpdate))

// 	})

// }

func UpdateCustomer(db *pg.DB, r *gin.Engine) {
	r.PUT("/customer/update", func(c *gin.Context) {
		addressUpdate := &dbpkg.AddressUpdate{}
		phoneUpdate := &dbpkg.PhoneUpdate{}
		if err := c.ShouldBindJSON(&phoneUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			fmt.Println(phoneUpdate)
			fmt.Println(phoneUpdate.CONTACT)
			fmt.Println(phoneUpdate.ID)

			c.JSON(http.StatusOK, dbpkg.CustomerUpdate(db, *phoneUpdate))
		}

		if err := c.ShouldBindJSON(&addressUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			fmt.Println(addressUpdate)
			fmt.Println(addressUpdate.ADDRESS)
			fmt.Println(addressUpdate.ID)
			c.JSON(http.StatusOK, dbpkg.CustomerUpdate(db, *addressUpdate))

		}
		// fmt.Println(phoneUpdate)
		// fmt.Println(&phoneUpdate.CONTACT)
		// fmt.Println(phoneUpdate.ID)

	})

}
