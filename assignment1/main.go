package main

import "fmt"

func main() {
	arr := []int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(sum(arr, 0))
}

func sum(arr []int, index int) int {
	if len(arr) <= index {
		return 0
	}
	return arr[index] + sum(arr, index+1)
}
