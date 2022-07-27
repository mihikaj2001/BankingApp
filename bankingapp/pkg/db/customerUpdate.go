package db

import (
	"bank_app/pkg/models"
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

type PhoneUpdate struct {
	ID      int64 `json:"customer_id" binding:"required"`
	CONTACT int64 `json:"contact_number" binding:"required"`
}

type AddressUpdate struct {
	ID      int64  `json:"customer_id" binding:"required"`
	ADDRESS string `json:"addr" binding:"required"`
}

// func CustomerUpdate(db *pg.DB, newUpdate *PhoneUpdate) models.Customers {
// 	var customer models.Customers
// 	fmt.Println(newUpdate.CONTACT)
// 	fmt.Println(newUpdate.ID)
// 	_, err := db.Model(&customer).Set("contact_number = ?", newUpdate.CONTACT).Where("customer_id = ?", newUpdate.ID).Update()
// 	fmt.Printf("err %#v", err)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return customer
// }

func CustomerUpdate(db *pg.DB, item interface{}) models.Customers {
	var customer models.Customers
	// fmt.Println(item.(type))
	x := fmt.Sprintf("%T", item)
	fmt.Println(x)
	switch item.(type) {
	case PhoneUpdate:
		newUpdate := item.(PhoneUpdate)
		// print(newUpdate)
		fmt.Println(newUpdate)
		fmt.Println(newUpdate.CONTACT)
		fmt.Println(newUpdate.ID)
		_, err := db.Model(&customer).Set("contact_number = ?", &newUpdate.CONTACT).Where("customer_id = ?", newUpdate.ID).Update()
		fmt.Printf("err %#v", err)
		if err != nil {
			panic(err)
		}

	case AddressUpdate:
		newUpdate := item.(AddressUpdate)
		// print(newUpdate)
		fmt.Println(newUpdate)
		fmt.Println(newUpdate.ADDRESS)
		fmt.Println(newUpdate.ID)
		_, err := db.Model(&customer).Set("addr = ?", &newUpdate.ADDRESS).Where("customer_id = ?", newUpdate.ID).Update()
		fmt.Printf("err %#v", err)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("error")

	}
	// fmt.Println(newUpdate.CONTACT)
	// fmt.Println(newUpdate.ID)
	// _, err := db.Model(&customer).Set("contact_number = ?", &newUpdate.CONTACT).Where("customer_id = ?", newUpdate.ID).Update()
	// fmt.Printf("err %#v", err)
	// if err != nil {
	// 	panic(err)
	// }
	return customer
}
