package main

import "testing"

var wordsToTest = []string{"baking", "biking", "baking cookies", "breaking", "king", "blinking eyes", "kingofblink"}
var searchWord = "blinking"

func BenchmarkLevenshtein1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			LevenshteinDistance(searchWord, word)
		}
	}
}

func BenchmarkLevenshtein2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			LevenshteinDistance2(searchWord, word)
		}
	}
}

func BenchmarkLevenshtein3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			LevenshteinDistance3(searchWord, word)
		}
	}
}

func BenchmarkLevenshtein4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			LevenshteinDistance4(&searchWord, &word)
		}
	}
}

func BenchmarkLevenshtein5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, word := range wordsToTest {
			LevenshteinDistance5(&searchWord, &word)
		}
	}
}
