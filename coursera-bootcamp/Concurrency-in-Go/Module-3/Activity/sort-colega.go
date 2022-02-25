package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortArr(arr []int, ch chan []int) {
	fmt.Println(arr)
	sort.Ints(arr[:])
	ch <- arr[:]
}

func main() {
	var arrInts = []int{}
	var ch = make(chan []int)

	// prompt inputs array of integers
	// example: 3 1 2 4 7 8 5 6 4 3 14 11
	// example paritition size: 3
	fmt.Println("Please input array of integers: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seriesOfInts := strings.Split(scanner.Text(), " ")

		// convert prompt inputs to int slice
		arrInts = make([]int, len(seriesOfInts))
		for i, v := range seriesOfInts {
			arrInts[i], _ = strconv.Atoi(v)
		}

		// split parition array
		partitionLen := len(arrInts) / 4

		// fmt.Println(arrInts[:partitionLen])

		// goroutine 1
		go sortArr(arrInts[:partitionLen], ch)
		// goroutine 2
		go sortArr(arrInts[partitionLen:(2 * partitionLen)], ch)
		// goroutine 3
		go sortArr(arrInts[(2 * partitionLen):(3 * partitionLen)], ch)
		// goroutine 4
		go sortArr(arrInts[(3 * partitionLen):], ch)

		mergeArr := []int{}
		mergeArr = append(mergeArr, <-ch...)
		mergeArr = append(mergeArr, <-ch...)
		mergeArr = append(mergeArr, <-ch...)
		mergeArr = append(mergeArr, <-ch...)

		// print merge array
		fmt.Println(mergeArr)
	}
}
