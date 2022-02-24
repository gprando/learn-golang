// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which
// is sorted by a different goroutine. Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼
// of the array should print the subarray that it will sort. When sorting is complete,
// the main goroutine should print the entire sorted list.

package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

var c = make(chan []int, 4)

func main() {

	var quantity int
	var inputtedNumber int
	var wg sync.WaitGroup
	sliceOfInts := make([]int, 0, 3)

	fmt.Println("Enter a quantity of number your like in array ?")
	fmt.Scan(&quantity)

	fmt.Println("Enter a series of integers to place into the array.")
	for i := 0; i < quantity; i++ {
		_, err := fmt.Scan(&inputtedNumber)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Invalid user input")
		}

		sliceOfInts = append(sliceOfInts, inputtedNumber)

	}
	fmt.Println("elements before sort", sliceOfInts)

	// fmt.Println("Here are your arrays when partitioned", slice1, slice2, slice3, slice4)

	sliceSize := quantity / 4

	fmt.Println("size of sub slices ", sliceSize)

	for i := 0; i < 4; i++ {
		var slice []int
		if i == 0 {
			slice = sliceOfInts[:sliceSize]
		} else if i == 3 {
			slice = sliceOfInts[3*(sliceSize):]
		} else {
			slice = sliceOfInts[i*sliceSize : i+1*sliceSize]
		}
		wg.Add(1)

		fmt.Println("slice", i, slice)
		go goroutineSort(slice, &wg)
	}
	wg.Wait()

	slice1 := <-c
	slice2 := <-c
	slice3 := <-c
	slice4 := <-c

	sliceSortred := mergeAndSort(slice1, slice2, slice3, slice4)

	fmt.Println("full slice sortred", sliceSortred)

}

func goroutineSort(unsortedSlice []int, wg *sync.WaitGroup) {
	sort.Ints(unsortedSlice)
	c <- unsortedSlice
	wg.Done()
}

func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
	fullSlice := mergeSlices(list1, list2, list3, list4)
	sort.Ints(fullSlice)
	return fullSlice
}

func mergeSlices(args ...[]int) []int {
	mergedSlice := make([]int, 0)
	for _, oneSlice := range args {
		mergedSlice = append(mergedSlice, oneSlice...)
	}

	return mergedSlice
}
