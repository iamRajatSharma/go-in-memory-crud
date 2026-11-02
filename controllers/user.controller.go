package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamrajatsharma/crud/models"
	"github.com/iamrajatsharma/crud/services"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is UP & Running....",
	})
}

func Users(c *gin.Context) {
	userList := services.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})
}

func GetUserById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	getUser, err := services.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": getUser,
	})

}

func AddUser(c *gin.Context) {
	var newUser models.User

	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := services.AddUser(newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func DeleteUserById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	deleteUser, err := services.DeleteUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": deleteUser,
	})

}

func UpdateUser(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	var newUser models.User

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	checkUser := c.ShouldBindJSON(&newUser)

	if checkUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": checkUser.Error(),
		})
		return
	}

	newUser, err1 := services.UpdateUser(newUser, id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err1,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User updated successfully",
		"data":    newUser,
	})

}
