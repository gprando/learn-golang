package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	var listofNames []Name
	fmt.Println("Please enter the name of the file")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
		names := strings.Split(scanner.Text(), " ")
		personName := Name{names[0], names[1]}
		listofNames = append(listofNames, personName)
	}

	file.Close()

	for i := 0; i < len(listofNames); i++ {
		fmt.Println(listofNames[i].fname + " " + listofNames[i].lname)
	}
}

type Name struct {
	fname, lname string
}
