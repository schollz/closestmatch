package main

// LevenshteinDistance computes the distance between two strings
//http://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)
	width := n - 1
	d := make([]int, m*n)

	for i := 1; i < m; i++ {
		d[i*width+0] = i
	}

	for j := 1; j < n; j++ {
		d[0*width+j] = j
	}

	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			if s[i] == t[j] {
				d[i*width+j] = d[(i-1)*width+(j-1)]
			} else {
				d[i*width+j] = Min(d[(i-1)*width+j]+1, //deletion
					d[i*width+(j-1)]+1,     //insertion
					d[(i-1)*width+(j-1)]+1) //substitution
			}
		}
	}
	return d[m*(width)+0]
}

// Min returns the minimal int among a list of integers
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// Alternative LevenshteinDistance (no min)
func LevenshteinDistance2(a string, b string) int {
	d := make([][]int, len(a)+1)
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, len(b)+1)
	}
	for i := 0; i < len(d); i++ {
		d[i][0] = i
	}
	for i := 0; i < len(d[0]); i++ {
		d[0][i] = i
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			ex := 1
			if a[i-1] == b[j-1] {
				ex = 0
			}
			min := d[i-1][j] + 1
			if (d[i][j-1] + 1) < min {
				min = d[i][j-1] + 1
			}
			if (d[i-1][j-1] + ex) < min {
				min = d[i-1][j-1] + ex
			}
			d[i][j] = min
		}
	}
	return d[len(a)][len(b)]
}

// Alternative LevenshteinDistance (no min, compute lengths once)
func LevenshteinDistance3(a string, b string) int {
	la := len(a)
	la1 := la + 1
	lb := len(b)
	lb1 := lb + 1

	d := make([][]int, la1)
	ld := len(d)
	for i := 0; i < ld; i++ {
		d[i] = make([]int, lb1)
	}
	for i := 0; i < ld; i++ {
		d[i][0] = i
	}
	ld0 := len(d[0])
	for i := 0; i < ld0; i++ {
		d[0][i] = i
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			ex := 1
			if a[i-1] == b[j-1] {
				ex = 0
			}
			min := d[i-1][j] + 1
			if (d[i][j-1] + 1) < min {
				min = d[i][j-1] + 1
			}
			if (d[i-1][j-1] + ex) < min {
				min = d[i-1][j-1] + ex
			}
			d[i][j] = min
		}
	}
	return d[la][lb]
}

// Alternative LevenshteinDistance (no min, compute lengths once, pointers)
func LevenshteinDistance4(a, b *string) int {
	la := len(*a)
	la1 := la + 1
	lb := len(*b)
	lb1 := lb + 1

	d := make([][]int, la1)
	ld := len(d)
	for i := 0; i < ld; i++ {
		d[i] = make([]int, lb1)
	}
	for i := 0; i < ld; i++ {
		d[i][0] = i
	}
	ld0 := len(d[0])
	for i := 0; i < ld0; i++ {
		d[0][i] = i
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			ex := 1
			if (*a)[i-1] == (*b)[j-1] {
				ex = 0
			}
			min := d[i-1][j] + 1
			if (d[i][j-1] + 1) < min {
				min = d[i][j-1] + 1
			}
			if (d[i-1][j-1] + ex) < min {
				min = d[i-1][j-1] + ex
			}
			d[i][j] = min
		}
	}
	return d[la][lb]
}

// Alternative LevenshteinDistance (no min, compute lengths once, pointers, 2 rows array)
func LevenshteinDistance5(a, b *string) int {
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
