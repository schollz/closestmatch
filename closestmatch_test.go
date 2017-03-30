package closestmatch

import (
	"fmt"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkClosest(b *testing.B) {
	cm := Open(test.WordsToTest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, searchWord := range test.SearchWords {
			cm.Closest(searchWord)
		}
	}
}

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Open(test.WordsToTest)
	}
}

func ExampleMatching() {
	cm := Open(test.WordsToTest)
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'don quixote by miguel de cervantes saavedra'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'charles dickens' matched 'hard times by charles dickens'
	// 'william shakespeare' matched 'the tragedy of romeo and juliet by william shakespeare'
	// 'war by hg wells' matched 'the war of the worlds by h. g. wells'
}
