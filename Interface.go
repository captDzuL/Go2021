package main

import (
	"fmt"
)

type hasName interface {
	GetName() string
}
type infKosong interface{}

func sayHello(has_Name hasName) {
	fmt.Println("Hello", has_Name.GetName())
}

func cobaInfKosong(infKosong interface{}) interface{} {
	if infKosong == true {
		return 1
	} else if infKosong == 2 {
		return "ups"
	} else {
		return false
	}
}

type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}
func main() {
	Cat := Animal{
		name: "Ushi",
	}

	sayHello(Cat)
	var data interface{} = cobaInfKosong(2)
	var data2 interface{} = cobaInfKosong(true)
	fmt.Println(data)
	fmt.Println(data2)
}
