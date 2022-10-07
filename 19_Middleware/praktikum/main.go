package main

import (
	"praktikum/controllers"
	"praktikum/database"
	"praktikum/middlewares"

	"github.com/labstack/echo/v4"
)

func init() {
	database.InitDB()
	database.InitialMigration()
}

func main() {
	// create a new echo instance
	e := echo.New()

	middlewares.LogMiddleware(e)

	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
