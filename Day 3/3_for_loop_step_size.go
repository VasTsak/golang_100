/*
The post iteration statement is called step size.
You can have a step size of 1 (increment by 1) or higher.
In this scenario we have a step size of 10.
*/

package main

import "fmt"

func main() {
	for i := 0; i < 101; i += 10 {
		fmt.Println(i)
	}
}
