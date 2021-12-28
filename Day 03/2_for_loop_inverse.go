/*
It is important to note, that a for loop can interate backwards.

All you need to do is to increment the post statement (3rd part)
*/

package main

import (
	"fmt"
)

func main() {
	for i := 25; i > 1; i-- {
		fmt.Println(i)
	}
}
