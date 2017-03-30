package main

import (
	"fmt"
	"testing"
)

func BenchmarkClosest(b *testing.B) {
	cm := Open(wordsToTest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, searchWord := range searchWords {
			cm.Closest(searchWord)
		}
	}
}

func TestGeneral(t *testing.T) {
	cm := Open(wordsToTest)
	fmt.Println("\nNew algorithm:\n")
	for _, searchWord := range searchWords {
		fmt.Printf("'%s'\tmatched\t'%s'\n", searchWord, cm.Closest(searchWord))
	}
}
