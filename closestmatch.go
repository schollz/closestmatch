package closestmatch

import (
	"math/rand"
	"sort"
	"strings"
)

type ClosestMatch struct {
	Substrings     map[string]map[string]struct{}
	SubstringSizes []int
}

func Open(possible []string, subsetSize []int) *ClosestMatch {
	cm := new(ClosestMatch)

	cm.SubstringSizes = subsetSize
	cm.Substrings = make(map[string]map[string]struct{})
	for _, s := range possible {
		s = strings.ToLower(s)
		cm.Substrings[s] = cm.splitWord(s)
	}

	return cm
}

func (cm *ClosestMatch) ClosestN(searchWord string, n int) []string {
	searchWordHash := cm.splitWord(searchWord)
	worstBestVal := 1000000
	bestWords := make(map[string]int)
	for word := range cm.Substrings {
		if len(bestWords) < n {
			newVal := cm.compareIfBetter(&searchWordHash, word, 0, len(word)+len(searchWord))
			bestWords[word] = newVal
			if newVal < worstBestVal {
				worstBestVal = newVal
			}
		} else {
			newVal := cm.compareIfBetter(&searchWordHash, word, worstBestVal, len(word)+len(searchWord))
			if newVal > worstBestVal {
				keyToDelete := ""
				newWorstBestVal := 100000
				for key, val := range bestWords {
					if val == worstBestVal {
						keyToDelete = key
					} else if val < newWorstBestVal {
						newWorstBestVal = val
					}
				}
				delete(bestWords, keyToDelete)
				bestWords[word] = newVal
				if newVal < newWorstBestVal {
					newWorstBestVal = newVal
				}
				worstBestVal = newWorstBestVal
			}

		}
	}

	// Return a sorted list
	bestWordsSlice := make([]string, len(bestWords))
	nm := map[int][]string{}
	var a []int
	for k, v := range bestWords {
		nm[v] = append(nm[v], k)
	}
	for k := range nm {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	i := 0
	for _, k := range a {
		for _, s := range nm[k] {
			bestWordsSlice[i] = s
			i++
		}
	}

	return bestWordsSlice[0:i]
}

func (cm *ClosestMatch) Closest(searchWord string) string {
	searchWordHash := cm.splitWord(searchWord)
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

func (cm *ClosestMatch) splitWord(word string) map[string]struct{} {
	wordHash := make(map[string]struct{})
	for _, j := range cm.SubstringSizes {
		mergedWord := word
		for it := 1; it < j; it++ {
			mergedWord = mergedWord + word
		}
		for i := 0; i < len(word); i++ {
			wordHash[string(mergedWord[i:i+j])] = struct{}{}
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

func (cm *ClosestMatch) Accuracy() float64 {
	rand.Seed(1)
	percentCorrect := 0.0
	numTrials := 0.0

	for wordTrials := 0; wordTrials < 100; wordTrials++ {

		var testString, originalTestString string
		testStringNum := rand.Intn(len(cm.Substrings))
		i := 0
		for s := range cm.Substrings {
			i++
			if i != testStringNum {
				continue
			}
			originalTestString = s
			break
		}

		// remove a random word
		for trial := 0; trial < 4; trial++ {
			words := strings.Split(originalTestString, " ")
			deleteWordI := rand.Intn(len(words))
			words = append(words[:deleteWordI], words[deleteWordI+1:]...)
			testString = strings.Join(words, " ")
			if cm.Closest(testString) == originalTestString {
				percentCorrect += 1.0
			}
			numTrials += 1.0
		}

		// remove a random word and reverse
		for trial := 0; trial < 4; trial++ {
			a := strings.Split(originalTestString, " ")
			if len(a) > 1 {
				deleteWordI := rand.Intn(len(a))
				a = append(a[:deleteWordI], a[deleteWordI+1:]...)
				for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
					a[left], a[right] = a[right], a[left]
				}
			}
			testString = strings.Join(a, " ")
			if cm.Closest(testString) == originalTestString {
				percentCorrect += 1.0
			}
			numTrials += 1.0
		}

		// remove a random word and shuffle and replace random letter
		for trial := 0; trial < 4; trial++ {
			a := strings.Split(originalTestString, " ")
			if len(a) > 1 {
				deleteWordI := rand.Intn(len(a))
				a = append(a[:deleteWordI], a[deleteWordI+1:]...)
				for i := range a {
					j := rand.Intn(i + 1)
					a[i], a[j] = a[j], a[i]
				}
			}
			testString = strings.Join(a, " ")
			letters := "abcdefghijklmnopqrstuvwxyz"
			if len(testString) == 0 {
				continue
			}
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii+1:]
			ii = rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii+1:]
			if cm.Closest(testString) == originalTestString {
				percentCorrect += 1.0
			}
			numTrials += 1.0
		}

	}

	return 100.0 * percentCorrect / numTrials
}
