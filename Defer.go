package main

import "fmt"

func logging() {
	fmt.Println("Ini Logging")
}

func runApp(value int) {
	defer logging()
	fmt.Println("aplikasi berjalan")
	result := 10 / value
	fmt.Println(result)
}
func main() {
	runApp(0)
}
