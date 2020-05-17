package CopyThroughSerialization

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address3 struct {
	StreetAddress, City, Country string
}

type Person3 struct {
	Name string
	Address3 *Address3
}

func (p *Person3) DeepCopy() *Person3 {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))
	d := gob.NewDecoder(&b)
	result := Person3{}
	_ = d.Decode(&result)

	return &result
}


func main(){
	john := Person3{"John", &Address3{
		"12 22x",
		"Chicago",
		"US",
	}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address3.StreetAddress = "new address"

	fmt.Println(john, john.Address3)
	fmt.Println(jane, jane.Address3)
}