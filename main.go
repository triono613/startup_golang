package main

import (
	"bwa-startup/handler"
	"bwa-startup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("connected")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1/")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_cheker", userHandler.CheckEmailAvailability)
	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.ID = 5
	// userInput.Name = "test save from service"
	// userInput.Email = "cobaaja@gmail.com"
	// userInput.Occupation = "programmer"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	ID:   3,
	// 	Name: "test saja",
	// }
	// userRepository.Save(user)
	// userService.RegisterUser(userInput)

	// fmt.Println("connection to db is grood")
	// var users []user.User
	// length := len(users)

	// fmt.Print(length)
	// db.Find(&users)

	// length = len(users)
	// fmt.Print(length)

	// for _, user := range users {

	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)

	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()

}
