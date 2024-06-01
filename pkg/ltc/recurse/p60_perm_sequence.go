package recurse

import (
	"fmt"
)

/*
123  1
132  2 -- 123  -> 23 -->
213  3
231  4
312  5
321  6
*/
func getPermutation(n int, k int) string {
    digits := make([]int, n)
    for i := 0; i < n; i++ {
        digits[i]=(i+1)
    }
    // k-1 to get to 0th index
    return fn(digits, k-1)
}

func fn(digits []int,  k int) string {
    if len(digits) == 1 {
        return fmt.Sprintf("%d", digits[0])
    }
    n := len(digits)
    permsPerDigit := fact(n)/n

    digit := k / permsPerDigit
    rem := k % permsPerDigit

    // can we optimize this
    remDigits := make([]int, 0, n-1)
    remDigits = append(remDigits, digits[:digit]...)
    remDigits = append(remDigits, digits[digit+1:]...)

    return fmt.Sprintf("%d%s", digits[digit], fn(remDigits, rem))
}

func fact(n int) int {
    ans := 1
    for ;n > 0; n-- {
        ans = ans * n
    }
    return ans
}
