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

func TestGeneral(t *testing.T) {
	cm := Open(test.WordsToTest)
	fmt.Println("\nNew algorithm:\n")
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s'\tmatched\t'%s'\n", searchWord, cm.Closest(searchWord))
	}
}
