package main

import (
	"fmt"
	"math"
)

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Print("Enter a acceleration: ")
	fmt.Scanf("%f", &acceleration)

	fmt.Print("Enter a initial velocity: ")
	fmt.Scanf("%f", &initialVelocity)

	fmt.Print("Enter a initial displacement: ")
	fmt.Scanf("%f", &initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Print("Enter time to calculate displace: ")
	fmt.Scanf("%f", &time)

	fmt.Println(fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	// s = ½ a t2 + vot + so
	return func(f float64) float64 {
		return (acceleration*math.Pow(f, 2))/2 + initialVelocity*f + initialDisplacement
	}
}
