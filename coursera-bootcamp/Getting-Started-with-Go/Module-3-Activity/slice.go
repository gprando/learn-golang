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
	sliceNumbers := make([]int, 3)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter number OR press X for close the program :D ")
	var i int
	var inputEntry string

	for input.Scan() {
		inputEntry = input.Text()
		inputInt, err := strconv.Atoi(inputEntry)

		if strings.EqualFold(inputEntry, "X") {
			break
		}

		if err != nil {
			fmt.Println("Input does not number and X, error ->", err.Error())
			break
		}

		sliceNumbers[i] = inputInt
		i++

		if len(sliceNumbers) == i {
			sliceNumbers = append(sliceNumbers, i)
		}

		sortSlice(sliceNumbers)
		fmt.Println("sortred ", sliceNumbers)

		fmt.Println("Enter new number OR press X to break program.")

	}

}

func sortSlice(sliceNumbers []int) {
	sort.Slice(sliceNumbers, func(i, j int) bool {
		return sliceNumbers[i] < sliceNumbers[j]
	})
}
