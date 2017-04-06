package closestmatch

import (
	"compress/gzip"
	"encoding/json"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// Closest match is the structure that contains the
// substring sizes and carrys a map of the substrings for
// easy lookup
type ClosestMatch struct {
	Substrings     map[string]map[string]struct{}
	SubstringSizes []int
}

// New returns a new structure for performing closest matches
func New(possible []string, subsetSize []int) *ClosestMatch {
	cm := new(ClosestMatch)

	cm.SubstringSizes = subsetSize
	cm.Substrings = make(map[string]map[string]struct{})
	for _, s := range possible {
		cm.Substrings[s] = cm.splitWord(strings.ToLower(s))
	}

	return cm
}

// Load can load a previously saved ClosetMatch object from disk
func Load(filename string) (*ClosestMatch, error) {
	cm := new(ClosestMatch)

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return cm, err
	}

	w, err := gzip.NewReader(f)
	if err != nil {
		return cm, err
	}

	err = json.NewDecoder(w).Decode(&cm)
	return cm, err
}

// Save writes the current ClosestSave object as a gzipped JSON file
func (cm *ClosestMatch) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	w := gzip.NewWriter(f)
	defer w.Close()
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	return enc.Encode(cm)
}

// ClosestN searches for the `searchWord` and returns the `n` closest matches
// as a string slice
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

// Closest searches for the `searchWord` and returns the closest match
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
		for i := 0; i < len(word)-j; i++ {
			wordHash[string(word[i:i+j])] = struct{}{}
		}
	}
	return wordHash
}

func (cm *ClosestMatch) compareIfBetter(one *map[string]struct{}, substring string, minPercentage int, lenSum int) int {
	minPercentage = minPercentage * lenSum / (2 * 1000)
	shared := 0
	two := cm.Substrings[substring]
	lenTwo := len(two)
	if len(*one) < lenTwo {
		numberLeft := len(*one)
		for item := range *one {
			if _, ok := two[item]; ok {
				shared++
			} else if numberLeft+shared < minPercentage {
				return (2 * 1000) / lenSum * shared
			}
			numberLeft--
		}
	} else {
		numberLeft := lenTwo
		for item := range two {
			if _, ok := (*one)[item]; ok {
				shared++
			} else if numberLeft+shared < minPercentage {
				return (2 * 1000) / lenSum * shared
			}
			numberLeft--
		}
	}
	return (2 * 1000) / lenSum * shared
}

// Accuracy runs some basic tests against the wordlist to
// see how accurate this bag-of-characters method is against
// the target dataset
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
			if len(words) < 3 {
				continue
			}
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
			words := strings.Split(originalTestString, " ")
			if len(words) > 1 {
				deleteWordI := rand.Intn(len(words))
				words = append(words[:deleteWordI], words[deleteWordI+1:]...)
				for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
					words[left], words[right] = words[right], words[left]
				}
			} else {
				continue
			}
			testString = strings.Join(words, " ")
			if cm.Closest(testString) == originalTestString {
				percentCorrect += 1.0
			}
			numTrials += 1.0
		}

		// remove a random word and shuffle and replace random letter
		for trial := 0; trial < 4; trial++ {
			words := strings.Split(originalTestString, " ")
			if len(words) > 1 {
				deleteWordI := rand.Intn(len(words))
				words = append(words[:deleteWordI], words[deleteWordI+1:]...)
				for i := range words {
					j := rand.Intn(i + 1)
					words[i], words[j] = words[j], words[i]
				}
			}
			testString = strings.Join(words, " ")
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

		if cm.Closest(testString) == originalTestString {
			percentCorrect += 1.0
		}
		numTrials += 1.0

	}

	return 100.0 * percentCorrect / numTrials
}
