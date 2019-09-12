package main

import (
	"strconv"
	"testing"
)

var (
	p Person
)

func TestValidateBenar(t *testing.T) {
	p.Name = "Jon Snow"
	p.Gender = "Male"
	p.Age = 24
	t.Logf("Validating " + p.Name + ", " + p.Gender + ", " + strconv.Itoa(p.Age))
	if p.Validate() != nil {
		t.Errorf("SALAH! Seharusnya jika tidak ada yang salah, akan mengembalikan nil")
	}
}

func TestValidateSalah(t *testing.T) {
	p.Name = ""
	p.Gender = "Hybrid"
	p.Age = -1
	t.Logf("Validating " + p.Name + ", " + p.Gender + ", " + strconv.Itoa(p.Age))
	if p.Validate() == nil {
		if p.Name == "" {
			t.Errorf("SALAH! Serharusnya Name Tidak Boleh Kosong!")
		}

		if p.Gender != "Male" && p.Gender != "Female" {
			t.Errorf("SALAH! Serharusnya Gender hanya ada male atau female!")
		}

		if p.Age < 0 {
			t.Errorf("SALAH! Serharusnya Age Tidak Boleh negatif!")
		}
	}
}
