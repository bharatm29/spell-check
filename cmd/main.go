package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	spellcheck "spell-check/internals/spellCheck"
)

func main() {
	if _, err := os.Stat("words_dictionary.json"); err != nil {
		fmt.Println("Required files not found, downloading...")

		file_link := "https://raw.githubusercontent.com/dwyl/english-words/master/words_dictionary.json"
		cmd := exec.Command("wget", file_link)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error downloading word dictionary file: %s\n", err)
		}

		fmt.Println("Finished downloading necessary files")
	}

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
        fmt.Printf("Enter a word: ")
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
