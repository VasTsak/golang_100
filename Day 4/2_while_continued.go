/*
In this exercise we illustrate an example of while loop being useful.

Other examples can be when you have an API running or a Kafka stream open, or a ML model
that is being trained.

When you train a ML model you don't know how many iterations the model needs to run
in order to converge, so you set a condition that is more related to performance and less
related to the number of iterations.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := 0
	for num != 15 {
		num = rand.Intn(20)
		fmt.Println(num)
	}
}
