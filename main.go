package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamrajatsharma/crud/controllers"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

var users = []User{
	{Id: 1, Name: "Rajat", Age: 25},
	{Id: 2, Name: "Aarav", Age: 28},
	{Id: 3, Name: "Priya", Age: 22},
	{Id: 4, Name: "Ananya", Age: 30},
	{Id: 5, Name: "Rohan", Age: 27},
	{Id: 6, Name: "Sneha", Age: 24},
	{Id: 7, Name: "Vikram", Age: 33},
	{Id: 8, Name: "Neha", Age: 29},
	{Id: 9, Name: "Karan", Age: 31},
	{Id: 10, Name: "Pooja", Age: 26},
}

func main() {
	router := gin.Default()

	router.GET("/health", controllers.Health)
	router.GET("/users", controllers.Users)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.AddUser)
	router.DELETE("/users/:id", controllers.DeleteUserById)
	router.PATCH("/users/:id", controllers.UpdateUser)

	router.Run()
}
