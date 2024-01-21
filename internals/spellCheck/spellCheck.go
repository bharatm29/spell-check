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

		if editDist >= 0 && editDist <= 2 {
			similarWords = append(similarWords, dicWord)
			fmt.Println(dicWord)
		}
	}

	return similarWords
}

func getEditDistance(word string, dicWord string) int {
	m := len(dicWord)
	n := len(word)

	dp := make([][]int, n+1)

	for idx := range dp {
		dp[idx] = make([]int, m+1)
		for j := range dp[idx] {
			dp[idx][j] = -1
		}
	}

	return editDist(n, m, word, dicWord, dp)
}

func editDist(i int, j int, a string, b string, dp [][]int) int {
	if i == 0 && j == 0 {
		return 0
	}

	if i == 0 && j != 0 {
		return j
	}

	if i != 0 && j == 0 {
		return i
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if a[i-1] == b[j-1] {
		return editDist(i-1, j-1, a, b, dp)
	}

	insertDist := 1 + editDist(i, j-1, a, b, dp)
	deleteDist := 1 + editDist(i-1, j, a, b, dp)
	replaceDist := 1 + editDist(i-1, j-1, a, b, dp)

	dp[i][j] = min(insertDist, deleteDist, replaceDist)
	return dp[i][j]
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
