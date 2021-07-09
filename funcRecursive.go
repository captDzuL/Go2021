package main

import "fmt"

func faktorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func faktorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRecursive(value-1)
	}
}

func main() {

	result := faktorialLoop(5)
	fmt.Println(result)
	result = faktorialRecursive(5)
	fmt.Println(result)
}
