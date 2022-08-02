package str

func wordSubsets(words1 []string, words2 []string) []string {
	univ := make([]string, 0)

	freqarr := make([][]int, len(words2))
	for i, w2 := range words2 {
		freq := charfrequency(w2)
		freqarr[i] = freq
	}

	mergedfreq := mergeFreqSet(freqarr)

	for _, w1 := range words1 {
		cmpfreq := charfrequency(w1)

		if isSuperSet(cmpfreq, mergedfreq) {
			univ = append(univ, w1)
		}

	}

	return univ

}

// from all set, keep only the largest since that will decide.
func mergeFreqSet(freqset [][]int) []int {
	merged := make([]int, 26)
	for _, freq := range freqset {
		for i := 0; i < 26; i++ {
			merged[i] = max(merged[i], freq[i])
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isSuperSet(super, sub []int) bool {
	for i := 0; i < len(sub); i++ {
		if super[i] < sub[i] {
			return false
		}
	}
	return true
}

func charfrequency(s string) []int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}
	return freq
}
