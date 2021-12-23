package main

import (
	"fmt"
)

func is_positive(number int) bool {
	if number > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	pos := 1
	neg := -1
	is_pos_1 := is_positive(pos)
	is_pos_2 := is_positive(neg)
	fmt.Println(is_pos_1)
	fmt.Println(is_pos_2)
}
