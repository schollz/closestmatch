
# closestmatch :page_with_curl:

<img src="https://img.shields.io/badge/version-2.0.0-brightgreen.svg?style=flat-square" alt="Version">
<a href="https://travis-ci.org/schollz/closestmatch"><img src="https://img.shields.io/travis/schollz/closestmatch.svg?style=flat-square" alt="Build Status"></a>
<a href="http://gocover.io/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/coverage-98%25-brightgreen.svg?style=flat-square" alt="Code Coverage"></a>
<a href="https://godoc.org/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>

*closestmatch* is a simple and fast Go library for fuzzy matching an input string to a list of target strings. *closestmatch* is useful for handling input from a user where the input (which could be mispelled or out of order) needs to match a key in a database. *closestmatch* uses a [bag-of-words approach](https://en.wikipedia.org/wiki/Bag-of-words_model) to precompute character n-grams to represent each possible target string. The closest matches have highest overlap between the sets of n-grams. The precomputation scales well and is much faster and more accurate than Levenshtein for long strings.

Getting Started
===============

## Install

```
go get -u -v github.com/schollz/closestmatch
```

## Use 

####  Create a *closestmatch* object from a list words

```golang
// Take a slice of keys, say band names that are similar
// http://www.tonedeaf.com.au/412720/38-bands-annoyingly-similar-names.htm
wordsToTest := []string{"King Gizzard", "The Lizard Wizard", "Lizzard Wizzard"}

// Choose a set of bag sizes, more is more accurate but slower
bagSizes := []int{2}

// Create a closestmatch object
cm := closestmatch.New(wordsToTest, bagSizes)
```

#### Find the closest match, or find the *N* closest matches

```golang
fmt.Println(cm.Closest("kind gizard"))
// returns 'King Gizzard'

fmt.Println(cm.ClosestN("kind gizard",3))
// returns [King Gizzard Lizzard Wizzard The Lizard Wizard]
```

#### Calculate the accuracy

```golang
// Calculate accuracy
fmt.Println(cm.Accuracy())
// ~ 53 % (still way better than Levenshtein which hits 0% with this particular set)

// Improve accuracy by adding more bags
bagSizes = []int{2, 3, 4}
cm = closestmatch.New(wordsToTest, bagSizes)
fmt.Println(cm.Accuracy())
// accuracy improves to ~ 75 %
```

#### Save/Load

```golang
// Save your current calculated bags
cm.Save("closestmatches.gob")

// Open it again
cm2, _ := closestmatch.Load("closestmatches.gob")
fmt.Println(cm2.Closest("lizard wizard"))
// prints "The Lizard Wizard"
```

### Accuracy and Speed

*closestmatch* is more accurate than Levenshtein for long strings (like in the test corpus). If you run `go test` the tests will pass which validate that Levenshtein performs < 60% accuracy and *closestmatch* performs with > 98% accuracy. 

*closestmatch* is 10-12x faster than [a fast implementation of Levenshtein](https://groups.google.com/forum/#!topic/golang-nuts/YyH1f_qCZVc). Try it yourself with the benchmarks:

```bash
cd $GOPATH/src/github.com/schollz/closestmatch && go test -bench=. > closestmatch.bench
cd $GOPATH/src/github.com/schollz/closestmatch/levenshtein && go test -bench=. > levenshtein.bench
benchcmp levenshtein.bench ../closestmatch.bench
```

which gives something like

```bash
benchmark                 old ns/op     new ns/op     delta
BenchmarkNew-8            1.52          1739997       +114473386.84%
BenchmarkClosestOne-8     424671        33654         -92.08%
BenchmarkLargeFile-8      121750600     11784608      -90.32%
```

The `New()` function is so much faster in *levenshtein* because there is no precomputation needed (obviously).

## Todo

- [x] ClosestN(n int) returns closest n matches
- [x] Function to compare accuracy (for tests?)
- [x] Open should have []int{1,2,3} for the specified substructure lengths, compare different lengths
- [x] Save/Load for precomputation (change Open -> New)
- [x] Use more intuitive variable names + improve documentation
- [x] How does this relate to bag of words?
- [ ] Compare to agrep (write a utility)
