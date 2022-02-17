package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	var animals []Animal

	fmt.Println("Accept comands `query (animal name) (action of animal)` OR  `newanimal (animal name) (animal type)`")

	for {

		fmt.Print(">")
		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input := strings.TrimSuffix(userQuery, "\n")
		userInput := strings.Split(input, " ")

		if len(userInput) != 3 {
			fmt.Println("Error try again")
			continue
		}

		switch userInput[0] {
		case "newanimal":
			if userInput[2] == "cow" {
				animals = append(animals, cow{name: userInput[1]})
			} else if userInput[2] == "snake" {
				animals = append(animals, snake{name: userInput[1]})
			} else if userInput[2] == "bird" {
				animals = append(animals, bird{name: userInput[1]})
			} else {
				fmt.Println("Invalid Query 2")
				fmt.Println(animals)
				continue
			}
			fmt.Println("Created it!")
		case "query":

			for _, animal := range animals {
				if animal.getName() == userInput[1] {
					if userInput[2] == "move" {
						animal.Move()
					} else if userInput[2] == "eat" {
						animal.Eat()
					} else if userInput[2] == "speak" {
						animal.Speak()
					}
				}
			}

		default:
			fmt.Println("Error")
			continue

		}
	}

}

type cow struct {
	name string
}

type snake struct {
	name string
}

type bird struct {
	name string
}

func (c cow) getName() string {
	return c.name
}

func (s snake) getName() string {
	return s.name
}

func (b bird) getName() string {
	return b.name
}

func (c cow) Eat() {
	fmt.Printf("%s grass\n", c.name)
}

func (c cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c cow) Speak() {
	fmt.Printf("%s moos\n", c.name)
}

func (s snake) Eat() {
	fmt.Printf("%s mice\n", s.name)
}

func (s snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s snake) Speak() {
	fmt.Printf("%s hisses\n", s.name)
}

func (b bird) Eat() {
	fmt.Printf("%s worms\n", b.name)
}

func (b bird) Move() {
	fmt.Printf("%s fly\n", b.name)
}

func (b bird) Speak() {
	fmt.Printf("%s peeps\n", b.name)
}
