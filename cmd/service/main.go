package main

import(
	"github.com/labstack/echo/v4"
	"errors"
	handlers "github.com/hassan-Sabeh/payment_gateway/handlers"

)

type Users struct{
	users map[int]string
}

var(
	usersMap = map[int]string{
		123 : "Hassan",
		456 : "Sabeh",
	}
	users = Users{
		usersMap,
	}
)

func (u *Users) AddUserToDb(userId int, userName string) error {
	_, ok := u.users[userId]
	if ok {
		return errors.New("user exists already")
	}
	
	u.users[userId] = userName
	return nil
}





func main() {
	e := echo.New()
	e.GET("/user", handlers.GetUser)
	e.POST("/add-user", handlers.AddUser)
	e.Logger.Fatal(e.Start(":1234"))
}
