package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}

	return fn
}

func main() {
		var inputStr string
		scanner := bufio.NewScanner(os.Stdin)

		var a float64
		var v0 float64
		var s0 float64
		var t float64

		fmt.Println("Input acceleration(a) initial velocity(v0) initial displacement(s0)")
		if scanner.Scan() {
			inputStr = scanner.Text()
		}
		splitStr := strings.Split(inputStr, " ")

		a, _ = strconv.ParseFloat(splitStr[0], 64)
		v0, _ = strconv.ParseFloat(splitStr[1], 64)
		s0, _ = strconv.ParseFloat(splitStr[2], 64)

		fmt.Println("Input time(t)")
		fmt.Scan(&t)

		fn := GenDisplaceFn(a, v0, s0)

		fmt.Println(fn(t))

}
