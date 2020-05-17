package main

import "fmt"

type Person2 struct {
	Name string
	Age int
}

// factory function
//func NewPerson(name string, age int) Person {
//  return Person{name, age}
//}

func NewPerson2(name string, age int) *Person2 {
	return &Person2{name, age}
}

func main_() {
	// initialize directly
	p := Person2{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPerson2("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
