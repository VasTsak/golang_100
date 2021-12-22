/*
The step size can be negative and greater than 1.
The way to do it is by reducing the value the step size.
In this scenario we have a step size of -100.
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1000; i > -1; i += -100 {
		// it can also be i -= 100, but to me it looks nicer with a negative increment.
		fmt.Println(i)
	}
}
