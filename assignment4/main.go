package main

import (
	"errors"
	"fmt"
)

//Person struct
type Person struct {
	Name   string
	Gender string
	Age    int
}

//Validate function
func (p Person) Validate() error {
	if p.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if p.Gender != "Male" && p.Gender != "Female" {
		return errors.New("Gender is either Male or Female")
	}

	if p.Age < 0 {
		return errors.New("There is no such thing as negative age")
	}
	return nil
}

func main() {
	p := Person{}
	p.Name = "Jon Snow"
	p.Gender = "Male"
	p.Age = 24
	fmt.Println(p.Validate())
}
