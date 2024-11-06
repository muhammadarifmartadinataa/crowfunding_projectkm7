package main

import (
	"crowfundig/handler"
	"crowfundig/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@rif123#@tcp(127.0.0.1:3306)/crowfunding_miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewServide(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}

//input dari user
// handler mapping input dari user menjadi  struct input
// service mapping dari struct input  ke struct User
// repository save struct User ke db
