package levenshtein

import (
	"fmt"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Open(test.WordsToTest)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	cm := Open(test.WordsToTest)
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func ExampleMatching() {
	cm := Open(test.WordsToTest)
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'emma by jane austen'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'charles dickens' matched 'beowulf'
	// 'william shakespeare' matched 'the iliad by homer'
	// 'war by hg wells' matched 'beowulf'
}

func TestAccuray(t *testing.T) {
	cm := Open(test.WordsToTest)
	fmt.Println(cm.Accuracy())
	// Output:
	// [the war of the worlds by h. g. wells the time machine by h. g. wells the iliad by homer]
}
