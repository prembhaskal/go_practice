package str

func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	vmap := make(map[rune]bool)
	for _, rn := range vowels {
		vmap[rn] = true
	}

	sptr := 0
	eptr := len(s) - 1

	rns := []rune(s)

	for sptr < len(s) && eptr >= 0 && sptr < eptr {
		left := rns[sptr]
		if !vmap[left] {
			sptr++
			continue
		}

		right := rns[eptr]
		if !vmap[right] {
			eptr--
			continue
		}

		rns[sptr] = right
		rns[eptr] = left
		sptr++
		eptr--

	}

	return string(rns)
}
