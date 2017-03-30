package main

import (
	"strings"
)

type ClosestMatch struct {
	Substrings map[string]map[string]struct{}
}

func Open(possible []string) *ClosestMatch {
	cm := new(ClosestMatch)

	cm.Substrings = make(map[string]map[string]struct{})
	for _, s := range possible {
		s = strings.ToLower(s)
		cm.Substrings[s] = splitWord(s)
	}

	return cm
}

func (cm *ClosestMatch) Closest(searchWord string) string {
	searchWordHash := splitWord(searchWord)
	bestVal := 0
	bestWord := ""
	for word := range cm.Substrings {
		v := cm.Substrings[word]
		newVal := compareIfBetter(&searchWordHash, &v, bestVal, len(word)+len(searchWord))
		if newVal > bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}

func splitWord(word string) map[string]struct{} {
	wordHash := make(map[string]struct{})
	tripleWord := word + word
	for j := 2; j <= 3; j++ {
		for i := 0; i < len(word); i++ {
			wordHash[string(tripleWord[i:i+j])] = struct{}{}
		}
	}
	return wordHash
}

func compareIfBetter(one *map[string]struct{}, two *map[string]struct{}, minPercentage int, lenSum int) int {
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
