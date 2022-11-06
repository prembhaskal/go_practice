package bit

func toHex(num int) string {
	unum := uint(num)
	if num < 0 {
		unum = uint(1<<32 + num)
	}
	hex := make([]rune, 0)
	for {
		n1 := unum % 16
		hex = append(hex, getHexChar(n1))
		unum = unum / 16
		if unum <= 0 {
			break
		}
	}

	return string(reverse405(hex))
}

func reverse405(chrs []rune) []rune {
	i := 0
	j := len(chrs) - 1
	for i < j {
		chrs[i], chrs[j] = chrs[j], chrs[i]
		i++
		j--
	}
	return chrs
}

func getHexChar(n uint) rune {
	if n < 10 {
		return '0' + rune(n)
	}
	return 'a' + rune(n-10)
}
