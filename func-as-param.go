package main

import (
	"fmt"
	"strings"
)

func chatWithFilter(chat string, filter func(string) string) {
	chatFiltered := filter(chat)
	fmt.Println("Chat: ", chatFiltered)
}
func CensorWord(str string) string {

	censored := []string{"anjing", "shit"}

	var newSlice []string

	// check for empty slice
	if len(censored) <= 0 {
		return str
	}

	// convert str into a slice
	strSlice := strings.Fields(str)

	//check each words in strSlice against censored slice
	for position, word := range strSlice {
		for _, forbiddenWord := range censored {

			// NOTE : change between Index and EqualFold to see the different result

			if test := strings.Index(strings.ToLower(word), forbiddenWord); test > -1 {

				// calculate how many # for replacement
				replacement := strings.Repeat("*", len(word))

				strSlice[position] = replacement
				newSlice = append(strSlice[:position], strSlice[position:]...)
			}
		}
	}

	// convert []string slice back to string
	return strings.Join(newSlice, " ")

}
func main() {

	chatWithFilter("Anjing lo sial shit", CensorWord)
}
