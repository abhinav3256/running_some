package main

import "fmt"

func main() {
	var arr = []int{3, 1, 2, 10, 1}
	var sum int
	var result = []int{}

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]

		result = append(result, sum)

	}

	fmt.Println(result)
}
