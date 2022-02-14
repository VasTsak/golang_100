package main

import "fmt"

func findMax(arr []int) int {
	var temp int

	temp = arr[0]

	for _, e := range arr {
		if temp < e {
			temp = e
		}
	}

	return temp
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = 0
	}
	return a
}

func countingSort(arr []int) []int {
	// generate array from min to max
	//todo: wie unten lösen
	counter := makeRange(0, findMax(arr))

	// count
	for _, e := range arr {
		counter[e] += 1
	}

	// add prev val to curr
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	// copy to correct pos
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		e := arr[i]         // elem to add
		t := counter[e] - 1 // target pos

		res[t] = e
		counter[e] = counter[e] - 1
	}

	return res
}

func main() {
	data := []int{5, 1, 9, 51, 9, 2}

	fmt.Println("data")
	fmt.Println(data)

	sorted := countingSort(data)

	fmt.Println("sorted")
	fmt.Println(sorted)
}
