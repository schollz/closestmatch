package main

import "fmt"

func main() {
	words := []string{"the sword in the stone by t.h. white", "the once and future king by t.h. white", "the book of merlyn by t.h. white", "lord of the rings by j.r.r. tolkein", "lord of the rings: return of the king by j.r.r. tolkein"}
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range words {
		wordHashes[word] = hashWord(word)

	}

	searchWord := "jr.r. tolkein return of the king"
	searchWordHash := hashWord(searchWord)
	maxVal := 0
	fmt.Println(searchWord)
	for word := range wordHashes {
		x := compareHash(searchWordHash, wordHashes[word])
		len1 := len(searchWordHash)
		len2 := len(wordHashes[word])
		curVal := 2 * 1000 * x / (len1 + len2)
		if curVal > maxVal {
			maxVal = curVal
		}
		fmt.Println(word, x, len1, len2, curVal, LevenshteinDistance5(&searchWord, &word))
	}
}

func compareHash(one map[string]struct{}, two map[string]struct{}) int {
	oneInTwo := 0
	for item := range one {
		if _, ok := two[item]; ok {
			oneInTwo++
		}
	}
	return oneInTwo
}

func hashWord(word string) map[string]struct{} {
	wordHash := make(map[string]struct{})
	tripleWord := word + word + word
	for j := 1; j <= 3; j++ {
		for i := 0; i < len(word); i++ {
			wordHash[string(tripleWord[i:i+j])] = struct{}{}
		}
	}
	return wordHash
}
