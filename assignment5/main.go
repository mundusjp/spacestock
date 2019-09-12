package main

import "fmt"

//Case struct
type Case struct {
	Order int
	Name  string
}

var cases []Case

// change order value
func (pointerToCase *Case) reorder(i int) {
	(*pointerToCase).Order = i
}

func rearrange(cases []Case, from int, to int) []Case {
	cases[from-1], cases[to-1] = cases[to-1], cases[from-1]
	var newCases []Case
	// reorder the case
	for i, c := range cases {
		c := &c
		c.reorder(i + 1)
		newCases = append(newCases, *c)
	}
	fmt.Printf("%+v\n", newCases)
	return newCases
}

func main() {
	cases = []Case{
		{Order: 1, Name: "one"},
		{Order: 2, Name: "two"},
		{Order: 3, Name: "tri"},
		{Order: 4, Name: "fou"},
		{Order: 5, Name: "fiv"},
	}
	cases = rearrange(cases, 4, 2)
	cases = rearrange(cases, 1, 2)
	fmt.Printf("%+v\n", cases)
}
