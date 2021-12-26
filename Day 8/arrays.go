/*
The type [n]T is an array of n values of type T.
vbb
The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry;
Go provides a convenient way of working with arrays.
*/

package main

import (
	"fmt"
)

var arr [2]string

func main() {
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr)
}
