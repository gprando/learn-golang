package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	intSlice := make([]int, 3)

	input := bufio.NewScanner(os.Stdin)

	var userInputStr string

	fmt.Print("Enter a number OR type X to exit the program :- ")
	var index int
	for input.Scan() {

		userInputStr = input.Text()
		userInputInt, err := strconv.Atoi(userInputStr)

		if strings.EqualFold(userInputStr, "X") {
			fmt.Println("Input is X - existing the program.")
			break
		} else if err != nil {
			fmt.Println("Input is not a valid number: ", err.Error())
			break
		} else {

			intSlice[index] = userInputInt
			index++
			fmt.Println("index and len of slice->", index, len(intSlice))
			if len(intSlice) == index {
				intSlice = append(intSlice, index)
				fmt.Println("slice appedned ->", intSlice, len(intSlice), cap(intSlice))
			}
			sort.Ints(intSlice)
			fmt.Println("sorted slice is->", intSlice)
		}

		fmt.Println("enter a number OR type X to exit the program.")

	}

}
