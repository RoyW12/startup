package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) 

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println("database connected")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router :=gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users",userHandler.RegisterUser)
	api.POST("/sessions",userHandler.Login)
	api.POST("/email_checker",userHandler.CheckEmailAvailability)
	router.Run()

	
	//handler, mapping input dari user => struct input
	//input dari user 
	//service :melakukan mapping dari struct input ke struct User
	//repository
	//db

		/* tes aja*/
	// input := user.LoginUserInput{
	// 	Email: "berta.cm",
	// 	Password: "berta12",
	// }

	// user,err := userService.Login(input)
	// if err != nil{
	// 	fmt.Println("error mas")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)



}

