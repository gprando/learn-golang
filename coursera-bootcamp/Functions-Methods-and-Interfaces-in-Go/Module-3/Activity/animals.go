package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.sound)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Printf(">")
		scanner.Scan()
		userInput := scanner.Text()

		strs := strings.Split(userInput, " ")
		if len(strs) != 2 {
			fmt.Println("Enter two strings")
			continue
		}

		var animal Animal
		switch strs[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Printf("Animal type %s is invalid\n", strs[0])
		}

		switch strs[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("Action %s is invalid\n", strs[1])
		}
	}
}
