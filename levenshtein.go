package main

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

func FindBestLevenshtein(searchWord string, wordsToTest []string) string {
	bestVal := 10000
	bestWord := ""
	for _, word := range wordsToTest {
		newVal := LevenshteinDistance(&searchWord, &word)
		if newVal < bestVal {
			bestVal = newVal
			bestWord = word
		}
	}
	return bestWord
}
