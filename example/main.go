package main

import (
	"casino_royal/vault/api"
	"context"
	"fmt"
)

func main() {
	fmt.Println("calling user service")

	users, err := fetchUsers()

	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}

	fmt.Println(len(users))

	for _, user := range users {
		fmt.Printf("%+v \n", user)
	}
}

func fetchUsers() ([]*api.User, error) {
	api := api.NewClient(nil)
	users, _, err := api.Users.ListAll(context.Background())
	return users, err
}
