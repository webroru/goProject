package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string
	users := make(map[string]string)

	fmt.Println("Enter a name")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	users["name"] = input

	fmt.Println("Enter an address")
	_, err = fmt.Scan(&input)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	users["address"] = input

	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Serialization error:", err)
		return
	}
	fmt.Printf("Json: %s\n", usersJson)
}
