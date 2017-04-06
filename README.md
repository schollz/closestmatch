
<p align="center">
<img
    src="logo.png"
    width="260" height="80" border="0" alt="closestmatch">
<br>
<a href="https://travis-ci.org/schollz/closestmatch"><img src="https://img.shields.io/travis/schollz/closestmatch.svg?style=flat-square" alt="Build Status"></a>
<a href="http://gocover.io/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/coverage-98%25-brightgreen.svg?style=flat-square" alt="Code Coverage"></a>
<a href="https://godoc.org/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>
</p>

<p align="center">Fuzzy match a set of strings</a></p>

*closestmatch* is a simple library for finding the closest match in a set of strings. This is useful if you want the user to input a string and then find a key in a database that is closest to the user input.

## How does this work?

*closestmatch* uses a [bag-of-characters approach](https://en.wikipedia.org/wiki/Bag-of-words_model) to represent each possible key and then matches with the key that has the highest overlap between the sets.

Since it uses a map for the lookup table it is fast and requires very little memory and scales well for large datasets.

Getting Started
===============

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
// ~ 53 % (still way better than Levenshtein which hits 0%)

// Improve accuracy by adding more bags
bagSizes = []int{2, 3, 4}
cm = closestmatch.New(wordsToTest, bagSizes)
fmt.Println(cm.Accuracy())
// accuracy improves to ~ 75 %
```

#### Save/Load

```golang
// Save your current calculated bags
cm.Save("closestmatches.json")

// Open it again
cm2, _ := closestmatch.Load("closestmatches.json")
fmt.Println(cm2.Closest("lizard wizard"))
// prints "The Lizard Wizard"
```

### Accuracy and Speed

*closestmatch* is about 2x more accurate than Levenshtein for long strings (like in the test corpus). If you run `go test` the tests will pass which validate that Levenshtein performs < 60% accuracy and *closestmatch* performs with > 98% accuracy.

*closestmatch* is 6-7x faster than [a fast implementation of Levenshtein](https://groups.google.com/forum/#!topic/golang-nuts/YyH1f_qCZVc). Try it yourself with the benchmarks:

```bash
cd $GOPATH/src/github.com/schollz/closestmatch && go test -bench=. > closestmatch.bench
cd $GOPATH/src/github.com/schollz/closestmatch/levenshtein && go test -bench=. > levenshtein.bench
benchcmp levenshtein.bench ../closestmatch.bench
```

which gives something like

```bash
benchmark                 old ns/op     new ns/op     delta
BenchmarkNew-8            1.49          624681        +41924799.33%
BenchmarkClosestOne-8     432350        61401         -85.80%
BenchmarkLargeFile-8      122050000     19925964      -83.67%
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
