package str

func isAnagram(s string, t string) bool {
	freq1 := calcLetterFrequency(s)
	freq2 := calcLetterFrequency(t)

	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}

func calcLetterFrequency(s string) []int {
	freq := make([]int, 26)

	for _, ch := range s {
		freq[ch-'a']++
	}

	return freq
}
