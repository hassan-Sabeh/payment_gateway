package main

import (
	bank "github.com/hassan-Sabeh/payment_gateway/bank_gateway_simulation"
	db "github.com/hassan-Sabeh/payment_gateway/database"
	handlers "github.com/hassan-Sabeh/payment_gateway/handlers"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware" //Uncomment for logging
)

func main() {
	e := echo.New()
	//init the db connection object
	db := db.NewDbConnection()

	//init bank simulation object
	b := new(bank.BankGateway)

	// e.Use(middleware.Logger()) //Uncomment for logging

	//Set the database and bank simulation objects in the context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Set("db", db)
			c.Set("bank", b)
			return next(c)
		}
	})
	//Keeping routes here since there are only two
	e.GET("/payment/:id", handlers.GetPayment)
	e.POST("/process-payment", handlers.ProcessPayment)
	e.Logger.Fatal(e.Start(":1234"))
}
