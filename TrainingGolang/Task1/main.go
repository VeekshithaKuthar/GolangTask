package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0664)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data, err := os.ReadFile(string("data.txt"))
	if err != nil {
		fmt.Println(err)
	}
	//Find out number of words
	cleanWord := strings.ReplaceAll(string(data), "|", " ")
	words := strings.Fields(string(cleanWord))
	count := len(words)
	fmt.Printf("The total number of word %d in Text File\n", count)

	//Find out number word ocurrences
	WordMap := make(map[string]int)
	for _, word := range words {
		WordMap[word]++
	}
	fmt.Println(WordMap)

	//Find out the maximum word(s) ocurrence
	maxCount := 0
	maxWord := ""
	for word, count := range WordMap {
		if count > maxCount {
			maxCount = count
			maxWord = word

		}
	}
	fmt.Printf("The Word %s occured %d times.", maxWord, maxCount)

	//Concurrently use shared variable

}
