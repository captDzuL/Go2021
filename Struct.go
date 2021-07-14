package main

import "fmt"

type Customer struct { //like class object di java
	Name, Address string
	Age           int
}

func (customer Customer) sayHelloMethod(name string) {
	fmt.Println("Hello", name, "my name is ", customer.Name)
}

func main() {
	var Cust Customer

	Cust.Age = 20
	Cust.Name = "Dzul"
	fmt.Println(Cust.Age)
	fmt.Println(Cust.Name)

	Cust2 := Customer{
		Name:    "Dzull",
		Address: "Jaksel",
		Age:     23,
	}

	fmt.Println(Cust2)
	Cust2.sayHelloMethod("Boy")
}
