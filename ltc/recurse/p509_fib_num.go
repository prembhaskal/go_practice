package recurse

func fib(n int) int {
	// return fibiter(n)
	F := make(map[int]int)
	return fibrec(n, F)
}

func fibrec(n int, F map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n, ok := F[n]; ok {
		return n
	}

	F[n] = fibrec(n-1, F) + fibrec(n-2, F)
	return F[n]
}

func fibiter(n int) int {
	a := 0
	b := 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	for i := 2; i <= n; i++ {
		c := a + b

		a = b
		b = c
	}
	return b
}
