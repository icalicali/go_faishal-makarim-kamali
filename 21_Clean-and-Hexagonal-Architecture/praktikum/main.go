package main

import (
	"belajar-go-echo/config"
	controller "belajar-go-echo/controller/user"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))
	app.Start(":8080")
}
