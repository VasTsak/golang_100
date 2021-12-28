/*
While loops comes handy when you don't know how many times you want to loop for.

In go, you can write a while loop using a for loop by skipping the post statement.
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum <= 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
