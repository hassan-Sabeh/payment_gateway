package handlers

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"fmt"
)

func getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	user, ok := users.users[id]
	
	if !ok {
		return c.String(http.StatusNotFound, "user not found")
	} 
	return c.String(http.StatusOK, fmt.Sprintf("User is: %v", user))
}

func AddUser(c echo.Context) error {
	idRaw := c.QueryParam("id")
	userName := c.QueryParam("userName")
	id, convError := strconv.Atoi(idRaw)
	if convError != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	err := users.AddUserToDb(id, userName)
	if err != nil {
		return c.String(http.StatusForbidden, fmt.Sprintf("User already exists"))
	}
	return c.String(http.StatusOK, "user has been created")
}