package dp

// summary
// only consider 2 characters at a time, assign +1 and -1 to them
// run kadane's algo to find maxsum, with 1 condition that at least one -1 present in it.
func largestVariance(s string) int {
	freqmap := make(map[rune]int)
	for _, ch := range s {
		freqmap[ch]++
	}

	gmax := 0

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				continue
			}

			// assign char[i] to +1 and char[j] to -1
			// run Kadane's algo to find max
			ch1 := rune(i + 'a')
			ch2 := rune(j + 'a')

			if freqmap[ch1] == 0 || freqmap[ch2] == 0 {
				continue
			}

			rem_a := freqmap[ch1] // NOT needed
			rem_b := freqmap[ch2] // NOT needed

			has_a := false
			has_b := false

			mxsum := 0
			currsum := 0
			for _, ch := range s {
				val := 0
				if ch == ch1 {
					val = 1
					has_a = true
					rem_a--
				} else if ch == ch2 {
					val = -1
					has_b = true
					rem_b--
				}
				currsum = currsum + val

				if has_a && has_b && currsum > mxsum {
					mxsum = currsum
				}
				if currsum < 0 && rem_a > 0 && rem_b > 0 { // reset only when other a is available.
					has_b = false // for case "abbba" and 'abbbb'
					currsum = 0
				}
			}

			if mxsum > gmax {
				gmax = mxsum
				// fmt.Printf("ch1: %c, ch2: %c, gmax:%d\n", ch1, ch2, gmax)
			}
		}
	}
	return gmax
}
