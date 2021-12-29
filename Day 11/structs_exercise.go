package main

import (
	"fmt"
)

type Measurements struct {
	name   string
	height float32
	weight float32
}

func main() {

	vas := Measurements{"Vasilis", 182.0, 84.0}
	fmt.Printf("%s is %v cm tall and %v kg heavy", vas.name, vas.height, vas.weight)
}
