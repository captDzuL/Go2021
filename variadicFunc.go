package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10, 12, 10)
	fmt.Println(total)

	slice := []int{10, 20, 30, 42, 12}
	total = sumAll(slice...)
	fmt.Println(total)
}
