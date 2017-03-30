package closestmatch

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
		newVal := cm.compareIfBetter(&searchWordHash, word, bestVal, len(word)+len(searchWord))
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
	for j := 1; j <= 3; j++ {
		for i := 0; i < len(word); i++ {
			wordHash[string(tripleWord[i:i+j])] = struct{}{}
		}
	}
	return wordHash
}

func (cm *ClosestMatch) compareIfBetter(one *map[string]struct{}, substring string, minPercentage int, lenSum int) int {
	cons := 2 * 1000 / lenSum
	shared := 0
	two := cm.Substrings[substring]
	if len(*one) < len(two) {
		numberLeft := len(*one)
		for item := range *one {
			if _, ok := two[item]; ok {
				shared++
			} else if cons*(numberLeft+shared) < minPercentage {
				return cons * shared
			}
			numberLeft--
		}
	} else {
		numberLeft := len(two)
		for item := range two {
			if _, ok := (*one)[item]; ok {
				shared++
			} else if cons*(numberLeft+shared) < minPercentage {
				return cons * shared
			}
			numberLeft--
		}
	}
	return cons * shared
}
