package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodbye
	result := sayGoodBye("Dzul")
	fmt.Println(result)
	fmt.Println(getGoodbye("Dzulfa"))
}
