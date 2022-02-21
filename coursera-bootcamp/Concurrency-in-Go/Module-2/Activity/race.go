//race condition occurs because of concurrent access of two go-routines at same time.
//x = x+1 is not atomic instructions at machine code level. It comprises of LOAD X; MULTIPLY X; SAVE X.
// interleaving of these 3 instructions from two different go-routines at run-time is non-deterministic.
//final result vary acc. to interleaving, so it is not deterministic. It can output 8 or 10 or 40

package main

import (
	"fmt"
	"sync"
)

// its global
var x int

func multiply(wg *sync.WaitGroup, factor int) {
	defer wg.Done()
	x = x * factor
}

func main() {
	x = 2
	var wg sync.WaitGroup
	wg.Add(2)
	go multiply(&wg, 4)
	go multiply(&wg, 5)
	wg.Wait()
	fmt.Println(x)
}
