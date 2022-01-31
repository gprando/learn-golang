package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string

	fmt.Println("Enter fileName")
	fmt.Scanf("%s", &fileName)

	names := readNamesFromFile(fileName)

	for _, name := range names {
		fmt.Println(name.fname + " " + name.lname)
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
		if len(line) <= 0 || len(fullName) < 2 || err == io.EOF {
			break
		}

		names = append(names, Name{fullName[0], fullName[1]})
	}

	file.Close()

	return names
}
