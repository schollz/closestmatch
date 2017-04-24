package levenshtein

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(test.WordsToTest)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	bText, _ := ioutil.ReadFile("../test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest)
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func ExampleMatching() {
	cm := New(test.WordsToTest)
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'emma by jane austen'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'hard times by charles dickens' matched 'hard times by charles dickens'
	// 'complete william shakespeare' matched 'the iliad by homer'
	// 'war by hg wells' matched 'beowulf'

}

func TestAccuracyBookWords(t *testing.T) {
	bText, _ := ioutil.ReadFile("../test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest)
	accuracy := cm.AccuracyMutatingWords()
	fmt.Printf("Accuracy with mutating words in book list:\t%2.1f%%\n", accuracy)
}

func TestAccuracyBookletters(t *testing.T) {
	bText, _ := ioutil.ReadFile("../test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest)
	accuracy := cm.AccuracyMutatingLetters()
	fmt.Printf("Accuracy with mutating letters in book list:\t%2.1f%%\n", accuracy)
}

func TestAccuracyDictionaryletters(t *testing.T) {
	bText, _ := ioutil.ReadFile("../test/popular.txt")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest)
	accuracy := cm.AccuracyMutatingWords()
	fmt.Printf("Accuracy with mutating letters in dictionary:\t%2.1f%%\n", accuracy)
}
