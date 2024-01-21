package spellcheck

import (
	"sort"
	"sync"
)

type Word struct {
	word     string
	editDist int
}

func GetSimilarWords(word string, words []map[string]int, preferredDist int) []Word {
	similarWords := []Word{}

	var wg sync.WaitGroup

	for idx := range words {
		wg.Add(1)
		go findSimilarWords(word, preferredDist, &similarWords, words[idx], &wg)
	}

	wg.Wait()

	sort.Slice(similarWords, func(i, j int) bool {
		return similarWords[i].editDist > similarWords[j].editDist
	})

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

func findSimilarWords(word string, preferredDist int, similarWords *[]Word, wordDictonary map[string]int, wg *sync.WaitGroup) {
	for dicWord := range wordDictonary {
		editDist := getEditDistance(word, dicWord)

		if editDist >= 0 && editDist <= preferredDist {
			*similarWords = append(*similarWords, Word{editDist: editDist, word: dicWord})
		}
	}

	wg.Done()
}
