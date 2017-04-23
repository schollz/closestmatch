package closestmatch

import (
	"encoding/gob"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// ClosestMatch is the structure that contains the
// substring sizes and carrys a map of the substrings for
// easy lookup
type ClosestMatch struct {
	SubstringSizes []int
	SubstringToID  map[string]map[uint32]struct{}
	IDToKey        map[uint32]string
	NumSubstrings  map[uint32]int
}

// New returns a new structure for performing closest matches
func New(possible []string, subsetSize []int) *ClosestMatch {
	cm := new(ClosestMatch)

	cm.SubstringSizes = subsetSize
	cm.SubstringToID = make(map[string]map[uint32]struct{})
	cm.IDToKey = make(map[uint32]string)
	cm.NumSubstrings = make(map[uint32]int)
	for i, s := range possible {
		cm.IDToKey[uint32(i)] = s
		substrings := cm.splitWord(strings.ToLower(s))
		cm.NumSubstrings[uint32(i)] = len(substrings)
		for substring := range substrings {
			if _, ok := cm.SubstringToID[substring]; !ok {
				cm.SubstringToID[substring] = make(map[uint32]struct{})
			}
			cm.SubstringToID[substring][uint32(i)] = struct{}{}
		}
	}

	return cm
}

// Load can load a previously saved ClosestMatch object from disk
func Load(filename string) (*ClosestMatch, error) {
	cm := new(ClosestMatch)

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return cm, err
	}
	err = gob.NewDecoder(f).Decode(&cm)
	return cm, err
}

// Save writes the current ClosestSave object as a gzipped JSON file
func (cm *ClosestMatch) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(cm)
}

func (cm *ClosestMatch) match(searchWord string) map[string]int {
	searchSubstrings := cm.splitWord(searchWord)
	searchSubstringsLen := len(searchSubstrings)
	m := make(map[string]int)
	for substring := range searchSubstrings {
		if ids, ok := cm.SubstringToID[substring]; ok {
			for id := range ids {
				if _, ok2 := m[cm.IDToKey[id]]; !ok2 {
					m[cm.IDToKey[id]] = 0
				}
				m[cm.IDToKey[id]] += 200000 / (searchSubstringsLen + cm.NumSubstrings[id])
			}
		}
	}
	return m
}

// Closest searches for the `searchWord` and returns the closest match
func (cm *ClosestMatch) Closest(searchWord string) string {
	for _, pair := range rankByWordCount(cm.match(searchWord)) {
		return pair.Key
	}
	return ""
}

// ClosestN searches for the `searchWord` and returns the n closests matches
func (cm *ClosestMatch) ClosestN(searchWord string, n int) []string {
	matches := make([]string, n)
	j := 0
	for i, pair := range rankByWordCount(cm.match(searchWord)) {
		if i == n {
			break
		}
		matches[i] = pair.Key
		j = i
	}
	return matches[:j+1]
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (cm *ClosestMatch) splitWord(word string) map[string]struct{} {
	wordHash := make(map[string]struct{})
	for _, j := range cm.SubstringSizes {
		for i := 0; i < len(word)-j; i++ {
			wordHash[string(word[i:i+j])] = struct{}{}
		}
	}
	return wordHash
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
		testStringNum := rand.Intn(len(cm.IDToKey))
		i := 0
		for id := range cm.IDToKey {
			i++
			if i != testStringNum {
				continue
			}
			originalTestString = cm.IDToKey[id]
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

		// test the original string
		if cm.Closest(testString) == originalTestString {
			percentCorrect += 1.0
		}
		numTrials += 1.0

	}

	return 100.0 * percentCorrect / numTrials
}
