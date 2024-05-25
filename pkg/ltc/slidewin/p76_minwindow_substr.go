package slidewin

import "fmt"

// sliding window with matching using comparing unique counts (instead of map hence faster)
// TODO make it readable
func minWindow(s string, t string) string {
	// make map of t
	freq := make(map[byte]int)
	unique := 0
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
		if freq[t[i]] == 1 {
			unique++
		}
	}

	if len(t) > len(s) {
		return ""
	}

	// init slide window
	actual := 0
	slide := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		slide[s[i]]++
	}

	for k, v := range slide {
		if freq[k] > 0 && v >= freq[k] {
			actual++
		}
	}

	minlen := 1000000

	// track matching substring
	ms := 0
	me := -1

	// iterate and check
	cs := 0
	ce := len(t) - 1
	// fmt.Println("exp map")
	// printmap(freq)
	for ce < len(s) {
		// check current state and record if smallest
		// fmt.Println("actual map")
		// printmap(slide)
		// if isWindowSubstring(freq, slide) {
		// fmt.Printf("matching start cs: %d, ce: %d, unique: %d, actual: %d\n", cs, ce, unique, actual)
		if unique == actual {
			currlen := ce - cs + 1
			// fmt.Printf("matched map, cs: %d, ce: %d, currlen: %d, minlen: %d \n", cs, ce, currlen, minlen)
			if currlen < minlen {
				minlen = currlen
				ms = cs
				me = ce
			}
		}

		// inc end
		ce++
		if ce == len(s) {
			break
		}
		slide[s[ce]]++
		if freq[s[ce]] > 0 && slide[s[ce]] == freq[s[ce]] {
			actual++
		}
		// fmt.Printf("slide map ce: %d\n", ce)
		// printmap(slide)

		// remove from start if not matching or if already present
		for cs < ce {
			startch := s[cs]
			if freq[startch] == 0 {
				cs++
				slide[startch]--
			} else if slide[startch] > freq[startch] {
				cs++
				slide[startch]--
			} else {
				// fmt.Printf("breaking cs inc at cs: %d, s[cs]: %c, ce: %d, s[ce]: %c \n", cs, startch, ce, s[ce])
				// update actual
				break
			}
		}

		// fmt.Printf("matching cs: %d, ce: %d, unique: %d, actual: %d\n", cs, ce, unique, actual)

		if unique == actual {
			currlen := ce - cs + 1
			// fmt.Printf("matched map, cs: %d, ce: %d, currlen: %d, minlen: %d \n", cs, ce, currlen, minlen)
			if currlen < minlen {
				minlen = currlen
				ms = cs
				me = ce
			}
		}
	}

	// fmt.Printf("breaking at cs: %d, ce: %d\n", cs, ce)

	return s[ms : me+1]
}

// sliding window with matching using comparing maps
func minWindow1(s string, t string) string {
	// make map of t
	freq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	if len(t) > len(s) {
		return ""
	}

	// init slide window
	slide := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		slide[s[i]]++
	}

	minlen := 1000000

	// track matching substring
	ms := 0
	me := -1

	// iterate and check
	cs := 0
	ce := len(t) - 1
	// fmt.Println("exp map")
	// printmap(freq)
	for ce < len(s) {
		// check current state and record if smallest
		// fmt.Println("actual map")
		// printmap(slide)
		if isWindowSubstring(freq, slide) {
			currlen := ce - cs + 1
			// fmt.Printf("matched map, cs: %d, ce: %d, currlen: %d, minlen: %d \n", cs, ce, currlen, minlen)
			if currlen < minlen {
				minlen = currlen
				ms = cs
				me = ce
			}
		} else {
			// fmt.Println("not matched map")
		}

		// inc end
		ce++
		if ce == len(s) {
			break
		}
		slide[s[ce]]++
		// fmt.Printf("slide map ce: %d\n", ce)
		// printmap(slide)

		// remove from start if not matching or if already present
		for cs < ce {
			startch := s[cs]
			if freq[startch] == 0 {
				cs++
				slide[startch]--
			} else if slide[startch] > freq[startch] {
				cs++
				slide[startch]--
			} else {
				// fmt.Printf("breaking cs inc at cs: %d, s[cs]: %c, ce: %d, s[ce]: %c \n", cs, startch, ce, s[ce])
				break
			}
		}
	}

	// fmt.Printf("breaking at cs: %d, ce: %d\n", cs, ce)

	return s[ms : me+1]
}

func printmap(m map[byte]int) {
	fmt.Println("map: ")
	for k, v := range m {
		fmt.Printf("k: %c, v: %d ", k, v)
	}
	fmt.Println()
}

func isWindowSubstring(exp, act map[byte]int) bool {
	for k, v := range exp {
		if act[k] < v {
			return false
		}
	}
	return true
}
