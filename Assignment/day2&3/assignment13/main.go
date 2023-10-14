package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countWords(a string) int {
	if len(a) == 0 || strings.TrimSpace(a) == "" {
		return 0
	}

	words := strings.Fields(a)

	return len(words)
}

func main() {
	content, err := ioutil.ReadFile(`day2&3\assignment13\sample.txt`)
	if err != nil {
		fmt.Println("Error during reading the file", err)
		return
	}
	countOfWords := countWords(string(content))

	fmt.Println("Number of words in a file", countOfWords)
}
