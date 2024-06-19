package str

import (
	"fmt"
	"strings"
)

func computeLPS(pattern string) []int {
	m := len(pattern)

	// generate compute function
	// lps[x] = length of longest prefix, matching at suffix for array pattern[0,1,2,...x-1] = pattern[0:x]
	// longest proper prefix - LPS
	lps := make([]int, m)

	lps[0] = 0
	j := 0
	for q := 1; q < m; q++ {
		for ; j > 0 && pattern[j] != pattern[q]; {
			// means previously pattern[j-1] matched but next char did not match
			// so we try lps of pattern[0,1,2,...,j-1] next.
			j = lps[j-1] 
		}
		if pattern[j] == pattern[q] {
			j = j + 1
		}
		lps[q] = j
	}

	return lps
}

func kmpStringMatch(text, pattern string) bool {
	lps := computeLPS(pattern)
	m := len(pattern)
	// run kmp
	n := len(text)
	j := 0
	

	for i := 0; i < n; i++ {
		for ; j > 0 && pattern[j] != text[i]; {
			// means previously pattern[j-1] matched but next char did not match
			// so we try lps of pattern[0,1,2,...,j-1] next.
			j = lps[j-1]
		}
		if pattern[j] == text[i] {
			j = j+1
		}
		if j == m {
			fmt.Printf("matched at %d . i: %d\n", i - m + 1, i)
			return true
		}
	}

	return false
}

func builtInStringMatch(text, pattern string) bool {
	return strings.Contains(text, pattern)
}