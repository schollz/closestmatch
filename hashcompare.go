package main

import (
	"strings"

	"github.com/schollz/jsonstore"
)

func compareHash(one map[string]struct{}, two map[string]struct{}) int {
	oneInTwo := 0
	if len(one) < len(two) {
		for item := range one {
			if _, ok := two[item]; ok {
				oneInTwo++
			}
		}
	} else {
		for item := range two {
			if _, ok := one[item]; ok {
				oneInTwo++
			}
		}
	}

	return oneInTwo
}

func compareHashIfBetter(one *map[string]struct{}, two *map[string]struct{}, minPercentage int, lenSum int) int {
	cons := 2 * 1000 / lenSum
	oneInTwo := 0
	if len(*one) < len(*two) {
		numberLeft := len(*one)
		for item := range *one {
			if _, ok := (*two)[item]; ok {
				oneInTwo++
			} else if cons*(numberLeft+oneInTwo) < minPercentage {
				return cons * oneInTwo
			}
			numberLeft--
		}
	} else {
		numberLeft := len(*two)
		for item := range *two {
			if _, ok := (*one)[item]; ok {
				oneInTwo++
			} else if cons*(numberLeft+oneInTwo) < minPercentage {
				return cons * oneInTwo
			}
			numberLeft--
		}
	}
	return cons * oneInTwo
}

func hashWord(word string) map[string]struct{} {
	wordHash := make(map[string]struct{})
	tripleWord := word + word
	for j := 2; j <= 3; j++ {
		for i := 0; i < len(word); i++ {
			wordHash[string(tripleWord[i:i+j])] = struct{}{}
		}
	}
	return wordHash
}

func findBestWord(searchWord string, wordsToTest []string) string {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	searchWordHash := hashWord(searchWord)
	bestVal := 0
	bestWord := ""
	for _, word := range wordsToTest {
		foo := wordHashes[word]
		newVal := compareHashIfBetter(&searchWordHash, &foo, 400, len(word)+len(searchWord))
		if newVal > bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}

func findBestWordWithWordHash(searchWord string, wordsToTest []string, wordHashes map[string]map[string]struct{}) string {
	searchWordHash := hashWord(searchWord)
	bestVal := 0
	bestWord := ""
	for _, word := range wordsToTest {
		newVal := compareHash(searchWordHash, wordHashes[word])
		if newVal > bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}

func findBestWordWithWordHashIfBetter(searchWord string, wordsToTest []string, wordHashes map[string]map[string]struct{}) string {
	searchWordHash := hashWord(searchWord)
	bestVal := 0
	bestWord := ""
	for _, word := range wordsToTest {
		foo := wordHashes[word]
		newVal := compareHashIfBetter(&searchWordHash, &foo, bestVal, len(word)+len(searchWord))
		if newVal > bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}

type ClosestMatch struct {
	ks jsonstore.JSONStore
}

func Open(possible []string) (*ClosestMatch, error) {
	cm := new(ClosestMatch)

	for _, s := range possible {
		s = strings.ToLower(s)
		err := cm.ks.Set(s, hashWord(s))
		if err != nil {
			return cm, err
		}
	}

	return cm, nil
}

func (cm *ClosestMatch) Closest(searchWord string) string {
	searchWordHash := hashWord(searchWord)
	bestVal := 0
	bestWord := ""
	for _, word := range cm.ks.Keys() {
		var v map[string]struct{}
		cm.ks.Get(word, &v)
		newVal := compareHashIfBetter(&searchWordHash, &v, bestVal, len(word)+len(searchWord))
		if newVal > bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}
