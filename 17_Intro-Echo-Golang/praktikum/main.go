package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserEdit struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	userId := c.Param("id")

	id, _ := strconv.Atoi(userId)

	var foundUser User

	for _, user := range users {
		if user.Id == id {
			foundUser = user
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success get user",
		"users":    foundUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userId := c.Param("id")

	id, _ := strconv.Atoi(userId)

	var newUsers []User = []User{}

	for _, user := range users {
		if user.Id != id {
			newUsers = append(newUsers, user)
		}
	}

	users = newUsers

	return c.JSON(http.StatusOK, map[string]any{
		"messages": "user deleted",
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	userId := c.Param("id")

	id, _ := strconv.Atoi(userId)

	userEdit := UserEdit{}

	c.Bind(&userEdit)

	updatedUser := User{}

	for index, user := range users {
		if user.Id == id {
			user.Email = userEdit.Email
			user.Name = userEdit.Name
			user.Password = userEdit.Password
			users[index] = user
			updatedUser = users[index]
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"messages": "user deleted",
		"user":     updatedUser,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
