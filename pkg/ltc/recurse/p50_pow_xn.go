package recurse

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	if n == 0 {
		return 1
	}

	a := myPow(x, n/2)

	pow := a * a

	if n%2 == 0 {
		return pow
	}

	return pow * x
}
