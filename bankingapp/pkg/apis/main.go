package apis

import (
	"context"

	"bank_app/pkg/apis/customers"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func ConnectDB() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "root",
		Database: "bankdb",
	})
	defer db.Close()

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	StartServer(db)

	return
}

func StartServer(db *pg.DB) {
	r := gin.Default()
	// custId := 1
	// accountId := 2
	// customers.GetCustomerDetails(db, r, custId)
	// customers.GetCustomerAccountDetails(db, r, custId)
	// customers.GetCustomerLoanDetails(db, r, custId)
	// customers.GetCustomerTransactions(db, r, accountId)

	// customers.InsertCustomerDetails(db, r)
	customers.UpdateCustomerDetails(db, r)
	// customers.DeleteCustomerDetails(db, r, custId)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
