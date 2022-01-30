package wc_278

// TODO
// PENDING
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)
	// precalc := make([]int, n)

	// hash(s, p, m) = (val(s[0]) * p0 + val(s[1]) * p1 + ... + val(s[k-1]) * pk-1) mod m.
	for i := 1; i < n; i++ {
		// s_i := int(s[i]) - 97
		// p
	}
	return ""
}

func pow(n, p, MOD int64) int {
	var res int64
	res = 1
	for p > 0 {
		if (p & 1) == 1 {
			res = (res * n) % MOD
		}

		n = (n * n) % MOD
		p = p / 2
	}

	return int(res)
}
