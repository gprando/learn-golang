package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fName string
	lName string
}

func main() {
	var fileName string

	fmt.Println("Enter fileName")
	fmt.Scanf("%s", &fileName)

	names := readNamesFromFile(fileName)

	for index, name := range names {
		fmt.Println("index", index)
		fmt.Println("first name", name.fName)
		fmt.Println("last name", name.lName)
	}
}

func readNamesFromFile(fileName string) []Name {

	var names []Name

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		fullName := strings.Fields(line)
		names = append(names, Name{fullName[0], fullName[1]})

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return names
}
