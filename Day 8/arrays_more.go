package main

import (
	"fmt"
)

var arr [5]int

func main() {
	for i := 0; i < 5; i++ {
		arr[i] = i
		fmt.Println(arr)
	}
}
