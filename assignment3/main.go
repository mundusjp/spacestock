package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name   string
	gender string
	age    int
}

func NewPerson() *person {
	p := person{}
	return &p
}

func (p *person) Name(s string) *person {
	p.name = s
	return p
}

func (p *person) Gender(s string) *person {
	p.gender = s
	return p
}

func (p *person) Age(i int) *person {
	p.age = i
	return p
}

func main() {
	jon := &person{}
	jon = NewPerson().Name("Jon Snow").Gender("Male").Age(24)
	fmt.Printf("%v\n", jon.name+", "+jon.gender+", "+strconv.Itoa(jon.age))
}
