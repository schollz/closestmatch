package levenshtein

import (
	"math/rand"
	"strings"
)

// LevenshteinDistance
// from https://groups.google.com/forum/#!topic/golang-nuts/YyH1f_qCZVc
// (no min, compute lengths once, pointers, 2 rows array)
// fastest profiled
func LevenshteinDistance(a, b *string) int {
	la := len(*a)
	lb := len(*b)
	d := make([]int, la+1)
	var lastdiag, olddiag, temp int

	for i := 1; i <= la; i++ {
		d[i] = i
	}
	for i := 1; i <= lb; i++ {
		d[0] = i
		lastdiag = i - 1
		for j := 1; j <= la; j++ {
			olddiag = d[j]
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (*a)[j-1] == (*b)[i-1] {
				temp = 0
			} else {
				temp = 1
			}
			if (lastdiag + temp) < min {
				min = lastdiag + temp
			}
			d[j] = min
			lastdiag = olddiag
		}
	}
	return d[la]
}

type ClosestMatch struct {
	WordsToTest []string
}

func New(wordsToTest []string) *ClosestMatch {
	cm := new(ClosestMatch)
	cm.WordsToTest = wordsToTest
	return cm
}

func (cm *ClosestMatch) Closest(searchWord string) string {
	bestVal := 10000
	bestWord := ""
	for _, word := range cm.WordsToTest {
		newVal := LevenshteinDistance(&searchWord, &word)
		if newVal < bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}

func (cm *ClosestMatch) Accuracy() float64 {
	rand.Seed(1)
	percentCorrect := 0.0
	numTrials := 0.0

	for wordTrials := 0; wordTrials < 100; wordTrials++ {

		var testString, originalTestString string
		testStringNum := rand.Intn(len(cm.WordsToTest))
		i := 0
		for _, s := range cm.WordsToTest {
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

func (cm *ClosestMatch) AccuracySimple() float64 {
	rand.Seed(1)
	percentCorrect := 0.0
	numTrials := 0.0

	for wordTrials := 0; wordTrials < 500; wordTrials++ {

		var testString, originalTestString string
		testStringNum := rand.Intn(len(cm.WordsToTest))

		originalTestString = cm.WordsToTest[testStringNum]

		testString = originalTestString

		// letters to replace with
		letters := "abcdefghijklmnopqrstuvwxyz"

		choice := rand.Intn(3)
		if choice == 0 {
			// replace random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii+1:]
		} else if choice == 1 {
			// delete random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + testString[ii+1:]
		} else {
			// add random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii:]
		}
		closest := cm.Closest(testString)
		if closest == originalTestString {
			percentCorrect += 1.0
		} else {
			//fmt.Printf("Original: %s, Mutilated: %s, Match: %s\n", originalTestString, testString, closest)
		}
		numTrials += 1.0
	}

	return 100.0 * percentCorrect / numTrials
}

// AccuracyMutatingWords runs some basic tests against the wordlist to
// see how accurate this bag-of-characters method is against
// the target dataset
func (cm *ClosestMatch) AccuracyMutatingWords() float64 {
	rand.Seed(1)
	percentCorrect := 0.0
	numTrials := 0.0

	for wordTrials := 0; wordTrials < 200; wordTrials++ {

		var testString, originalTestString string
		testStringNum := rand.Intn(len(cm.WordsToTest))
		originalTestString = cm.WordsToTest[testStringNum]
		testString = originalTestString

		var words []string
		choice := rand.Intn(3)
		if choice == 0 {
			// remove a random word
			words = strings.Split(originalTestString, " ")
			if len(words) < 3 {
				continue
			}
			deleteWordI := rand.Intn(len(words))
			words = append(words[:deleteWordI], words[deleteWordI+1:]...)
			testString = strings.Join(words, " ")
		} else if choice == 1 {
			// remove a random word and reverse
			words = strings.Split(originalTestString, " ")
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
		} else {
			// remove a random word and shuffle and replace 2 random letters
			words = strings.Split(originalTestString, " ")
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
		}
		closest := cm.Closest(testString)
		if closest == originalTestString {
			percentCorrect += 1.0
		} else {
			//fmt.Printf("Original: %s, Mutilated: %s, Match: %s\n", originalTestString, testString, closest)
		}
		numTrials += 1.0
	}
	return 100.0 * percentCorrect / numTrials
}

// AccuracyMutatingLetters runs some basic tests against the wordlist to
// see how accurate this bag-of-characters method is against
// the target dataset when mutating individual letters (adding, removing, changing)
func (cm *ClosestMatch) AccuracyMutatingLetters() float64 {
	rand.Seed(1)
	percentCorrect := 0.0
	numTrials := 0.0

	for wordTrials := 0; wordTrials < 200; wordTrials++ {

		var testString, originalTestString string
		testStringNum := rand.Intn(len(cm.WordsToTest) - 1)
		originalTestString = cm.WordsToTest[testStringNum]
		testString = originalTestString

		// letters to replace with
		letters := "abcdefghijklmnopqrstuvwxyz"

		choice := rand.Intn(3)
		if choice == 0 {
			// replace random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii+1:]
		} else if choice == 1 {
			// delete random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + testString[ii+1:]
		} else {
			// add random letter
			ii := rand.Intn(len(testString))
			testString = testString[:ii] + string(letters[rand.Intn(len(letters))]) + testString[ii:]
		}
		closest := cm.Closest(testString)
		if closest == originalTestString {
			percentCorrect += 1.0
		} else {
			//fmt.Printf("Original: %s, Mutilated: %s, Match: %s\n", originalTestString, testString, closest)
		}
		numTrials += 1.0
	}

	return 100.0 * percentCorrect / numTrials
}
