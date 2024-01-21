package spellcheck

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetSimilarWords(word string) []string {
	wordDictonary := getWordDictonary()

	_, ok := wordDictonary[word]

	if ok {
		return []string{word}
	}

	similarWords := []string{}

	for dicWord := range wordDictonary {
		editDist := getEditDistance(word, dicWord)

		if editDist >= 0 && editDist <= 3 {
			similarWords = append(similarWords, dicWord)
            fmt.Println(dicWord)
		}
	}

	return similarWords
}

func getEditDistance(word string, dicWord string) int {
	return editDist(len(word), len(dicWord), word, dicWord)
}

func editDist(i int, j int, a string, b string) int {
	if i == 0 && j == 0 {
		return 0
	}

	if i == 0 && j != 0 {
		return j
	}

	if i != 0 && j == 0 {
		return i
	}

	if a[i-1] == b[j-1] {
		return editDist(i-1, j-1, a, b)
	}

	insertDist := 1 + editDist(i, j-1, a, b)
	deleteDist := 1 + editDist(i-1, j, a, b)
	replaceDist := 1 + editDist(i-1, j-1, a, b)

	return min(insertDist, deleteDist, replaceDist)
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
