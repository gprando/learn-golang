package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkTheString(userString string) bool {
	return strings.HasPrefix(userString, "i") &&
		strings.HasSuffix(userString, "n") &&
		strings.Contains(userString, "a")
}

func main() {
	fmt.Println("Enter a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userString := scanner.Text()
	fmt.Println(userString)

	if checkTheString(strings.ToLower(userString)) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
