package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string

	fmt.Println("Enter name")
	fmt.Scanf("%s", &name)

	fmt.Println("Enter address")
	fmt.Scanf("%s", &addr)

	personPropertiesMap := map[string]string{
		"name": name,
		"addr": addr,
	}

	jsonStringArray, err := json.Marshal(personPropertiesMap)

	if err != nil {
		fmt.Println("error in map convertion")
		return
	}

	fmt.Println("name and addr converted to json", string(jsonStringArray))
}
