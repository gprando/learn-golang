package main

import (
	"fmt"
	"strings"
)

var checkContais = [3]string{"i", "a", "n"}

func main() {
	fmt.Println("insert your string")

	var stringUser string

	fmt.Scanf("%s", &stringUser)

	initWithI := strings.ToLower(stringUser[0:1]) == "i"
	endWithN := strings.ToLower(stringUser[len(stringUser)-1:]) == "n"
	containsA := strings.Contains(strings.ToLower(stringUser), "a")

	if initWithI && containsA && endWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
