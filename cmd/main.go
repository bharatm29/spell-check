package main

import (
	"encoding/json"
	"fmt"
	"os"
	spellcheck "spell-check/internals/spellCheck"
)

func main() {
	var word string
	wordDictonary := getWordDictonary()

	words := []map[string]int{}

	cnt := 0
	limit := 50_000
	dicWords := map[string]int{}

	for key, val := range wordDictonary {
		dicWords[key] = val
		cnt++
		if cnt%limit == 0 {
			words = append(words, dicWords)
			dicWords = map[string]int{}
		}
	}

	for {
		fmt.Scanln(&word)

		_, ok := wordDictonary[word]

		if ok {
			fmt.Println([]string{word})
		} else {
			fmt.Println(spellcheck.GetSimilarWords(word, words, 2))
		}
	}
}

func getWordDictonary() map[string]int {
	b, err := os.ReadFile("words_dictionary.json")
	if err != nil {
		fmt.Print(err)
	}

	jsonStr := string(b)

	result := make(map[string]int)

	json.Unmarshal([]byte(jsonStr), &result)

	return result
}
