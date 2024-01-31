package main

import (
	"encoding/json"
	"fmt"
	"os"
	"your_project/model"
)

func main() {

	// read json file
	byteValue, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// put these users into temporary slice
	var temporaryUsers []model.User
	if err := json.Unmarshal(byteValue, &temporaryUsers); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	//create user with every user from temporary slice
	for _, tempUser := range temporaryUsers {
		user := model.CreateUser(tempUser.FirstName, tempUser.LastName)
		fmt.Printf("Added User: ID=%d, FirstName=%s, LastName=%s\n", user.ID, user.FirstName, user.LastName)
	}

	user, ok := model.GetUser(3)

	if ok {
		fmt.Println(user.FirstName + ` ` + user.LastName)
	} else {
		fmt.Println("User not found")
	}

	user, ok = model.GetUser(1)

	if ok {
		fmt.Println(user.FirstName + ` ` + user.LastName)
	} else {
		fmt.Println("User not found")
	}

}
