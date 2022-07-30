package db

import (
	"bank_app/pkg/models"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v10"
)

type PhoneUpdate struct {
	ID      int64 `json:"id"`
	CONTACT int64 `json:"contact_number"`
}

type AddressUpdate struct {
	ID      int64  `json:"id"`
	ADDRESS string `json:"addr"`
}

func CustomerUpdate(db *pg.DB, item interface{}) models.Customers {
	var customer models.Customers
	// fmt.Println(item.(type))
	x := fmt.Sprintf("%T", item)
	fmt.Println(x)
	switch item.(type) {
	case PhoneUpdate:
		newUpdate := item.(PhoneUpdate)
		// print(newUpdate)
		// fmt.Println(newUpdate)
		// fmt.Println(newUpdate.CONTACT)
		// fmt.Println(newUpdate.ID)
		_, err := db.Model(&customer).Set("contact_number = ?0, updated_at = ?1", &newUpdate.CONTACT, time.Now()).Where("id = ?", newUpdate.ID).Update()
		fmt.Printf("err %#v", err)
		if err != nil {
			panic(err)
		}

	case AddressUpdate:
		newUpdate := item.(AddressUpdate)
		// print(newUpdate)
		// fmt.Println(newUpdate)
		// fmt.Println(newUpdate.ADDRESS)
		// fmt.Println(newUpdate.ID)
		_, err := db.Model(&customer).Set("addr = ?0, updated_at = ?1", &newUpdate.ADDRESS, time.Now()).Where("id = ?", newUpdate.ID).Update()
		fmt.Printf("err %#v", err)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("error")

	}

	return customer
}
