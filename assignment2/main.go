package main

import "fmt"

func main() {
	arr := []string{
		"banana",
		"potato",
		"tomato",
		"manggo",
		"ulala",
	}
	word := "banana"
	fmt.Println(contains(arr, word, 0))
}

func contains(arr []string, word string, index int) bool {
	if len(arr) <= index {
		return false
	}
	if arr[index] == word {
		return true
	} else {
		return contains(arr, word, index+1)
	}

}
