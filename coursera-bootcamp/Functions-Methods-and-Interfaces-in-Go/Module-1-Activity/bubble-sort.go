package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	slice := rand.Perm(10)
	fmt.Println("all positive, before sort", slice)

	BubbleSort(slice)
	fmt.Println("all positive, after sort", slice)

	slice = rand.Perm(10)
	slice[0] = -4
	slice[4] = -10
	slice[7] = -8
	fmt.Println("with negative, before sort", slice)

	BubbleSort(slice)
	fmt.Println("with negative, after sort", slice)
}

func BubbleSort(sliceOfInt []int) {
	for i := 0; i < len(sliceOfInt)-1; i++ {
		for j := 0; j < len(sliceOfInt)-i-1; j++ {
			if sliceOfInt[j] > sliceOfInt[j+1] {
				Swap(sliceOfInt, j)
			}
		}
	}

}

func Swap(sliceOfInt []int, index int) {
	sliceOfInt[index], sliceOfInt[index+1] = sliceOfInt[index+1], sliceOfInt[index]
}
