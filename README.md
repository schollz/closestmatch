
<p align="center">
<img
    src="logo.png"
    width="260" height="80" border="0" alt="closestmatch">
<br>
<a href="https://travis-ci.org/schollz/closestmatch"><img src="https://img.shields.io/travis/schollz/closestmatch.svg?style=flat-square" alt="Build Status"></a>
<a href="http://gocover.io/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/coverage-0%25-red.svg?style=flat-square" alt="Code Coverage"></a>
<a href="https://godoc.org/github.com/schollz/closestmatch"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>
</p>

<p align="center">Get the closest match using substrings</a></p>

*closestmatch* tries to use a list of keywords to generate the closest match


Getting Started
===============

## Comparison with `agrep`

```
perf stat -r 50 -d agrep -iBy 'on one condition' test/books.list
```

## Todo

- [ ] Use more intuitive variable names + improve documentation
- [ ] ClosestN(n int) returns closest n matches
- [ ] Function to compare accuracy (for tests?)
- [ ] Open should have []int{1,2,3} for the specified substructure lengths, compare different lengths
- [ ] Save/Load for precomputation
- [ ] Compare to agrep (write a utility)
