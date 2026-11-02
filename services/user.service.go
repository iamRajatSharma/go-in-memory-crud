package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iamrajatsharma/crud/models"
)

var users = []models.User{
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

func GetAllUsers() []models.User {
	return users
}

func GetUserById(id int) (models.User, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func AddUser(user models.User) (models.User, error) {

	for _, userList := range users {
		if strings.EqualFold(userList.Name, user.Name) {
			return models.User{}, errors.New("User alread exists with this name")
		}
	}

	maxId := users[len(users)-1]

	user.Id = maxId.Id + 1
	users = append(users, user)
	return user, nil
}

func DeleteUserById(id int) (string, error) {
	getInder := -1
	for index, value := range users {
		if value.Id == id {
			getInder = index
			break
		}
	}

	if getInder == -1 {
		return "", errors.New("No user foudn with that id")
	}

	users = append(users[:getInder], users[getInder+1:]...)
	return "user deleted successfully", nil
}

func UpdateUser(newUser models.User, id int) (models.User, error) {
	fmt.Printf("%T", id)
	getuserId := -1
	for index, value := range users {
		if value.Id == id {
			getuserId = index
			break
		}
	}

	if getuserId == -1 {
		return models.User{}, errors.New("No user foudn with that id")
	}

	users[getuserId].Age = newUser.Age
	users[getuserId].Name = newUser.Name
	return users[id], nil
}
