/*
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

package main

import "fmt"

var c, python, java bool

func main() {
	k := 3
	var i, j int
	r, golang, cpp := true, false, "no!"

	fmt.Println(i, j, k, c, python, java, r, golang, cpp)
}
