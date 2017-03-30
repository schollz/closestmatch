package main

import "testing"

func BenchmarkSubDistance(b *testing.B) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	searchWordHash := hashWord(searchWord)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			compareHash(searchWordHash, wordHashes[word])
		}
	}
}
