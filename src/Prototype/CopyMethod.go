package Prototype

import "fmt"

type Address2 struct {
	StreetAddress, City, Country string
}

type Person2 struct {
	Name string
	Address2 *Address2
	Friends []string
}

func (a *Address2) DeepCopy() *Address2 {
	return &Address2 {
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func (p *Person2) DeepCopy() *Person2 {
	q := *p //we are making a copy of the actual thing
	q.Address2 = p.Address2.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main(){
	john := Person2{"John", &Address2{
		"123 rutger",
		"St Louis",
		"US",
	},
	[]string{"ay", "friends"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address2.StreetAddress = "230 main st"
	jane.Friends = append(jane.Friends, "jose")
	fmt.Println(jane)
	fmt.Println(john)

}