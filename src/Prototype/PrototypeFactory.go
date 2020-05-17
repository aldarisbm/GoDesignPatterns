package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address4 struct {
	Suite int
	StreetAddress, City string
}

type Employee struct {
	Name string
	Office Address4
}

func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}

	en := gob.NewEncoder(&b)
	_ = en.Encode(e)

	de := gob.NewDecoder(&b)
	result := Employee{}
	_ = de.Decode(&result)
	return &result
}

var mainOffice = Employee {
	"",
	Address4{0, "123 west drive", "London"},
}

var auxOffice = Employee {
	"",
	Address4{0, "33 east drive", "London"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 202)

	fmt.Println(john)
	fmt.Println(jane)
}
