package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceNumbers := []int{34, 22, 1}

	sortSlice(sliceNumbers)
	fmt.Println(sliceNumbers)

	sliceNumbers = append(sliceNumbers, 16)
	sortSlice(sliceNumbers)
	fmt.Println(sliceNumbers)
}

func sortSlice(sliceNumbers []int) {
	sort.Slice(sliceNumbers, func(i, j int) bool {
		return sliceNumbers[i] < sliceNumbers[j]
	})
}
