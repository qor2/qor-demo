package main

import (
	"github.com/qor/transition"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Order struct {
	ID uint
	transition.Transition
}

func main() {
	db, _ := gorm.Open("sqlite3", "demo.db")
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&Order{})
	db.AutoMigrate(&transition.StateChangeLog{})

	OrderStateMachine := transition.New(&Order{})
	OrderStateMachine.Initial("draft")
	OrderStateMachine.State("checkout")
	OrderStateMachine.State("paid").Enter(func(order interface{}, tx *gorm.DB) error {
		return nil
	}).Exit(func(order interface{}, tx *gorm.DB) error {
		return nil
	})

	OrderStateMachine.State("cancelled")
	OrderStateMachine.State("paid_cancelled")

	// define an event
	OrderStateMachine.Event("checkout").To("checkout").From("draft")

	OrderStateMachine.Event("paid").To("paid").From("checkout").Before(func(order interface{}, tx *gorm.DB) error {
		// business logic here
		fmt.Println("before paid")
		return nil
	}).After(func(order interface{}, tx *gorm.DB) error {
		// business logic here
		fmt.Println("after paid")
		return nil
	})

	cancelledEvent := OrderStateMachine.Event("cancel")
	cancelledEvent.To("cancelled").From("draft", "checkout")
	cancelledEvent.To("paid_cancelled").From("paid").After(func(order interface{}, tx *gorm.DB) error {
		// refund
		fmt.Println("refund cancel")
		return nil
	})

	order := Order{}

	OrderStateMachine.Trigger("checkout", &order, nil)
	OrderStateMachine.Trigger("paid", &order, db, "charged offline by jinzhu")
	OrderStateMachine.Trigger("cancel", &order, nil)

	fmt.Println(order.GetState())

}
