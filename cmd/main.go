package main

import (
	"fmt"
	spellcheck "spell-check/internals/spellCheck"
)

func main() {
	var word string

	fmt.Scanln(&word)

	fmt.Println(spellcheck.GetSimilarWords(word))
}
