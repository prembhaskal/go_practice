package recurse

func partition(s string) [][]string {
    ans_130 = make([][]string, 0)
    fnpart(0, len(s)-1, []rune(s), make([]string,0))
    return ans_130
}
var ans_130 [][]string

// push down answers
// collect all partitions at base case
func fnpart(start, end int, s []rune, prev []string) {
    if start > end {
        cp := make([]string, len(prev))
        copy(cp, prev)
        ans_130 = append(ans_130, cp)
        return
    }


    for i := start; i <= end; i++ {
        if isPalin(s, start, i) {
            prev = append(prev, string(s[start:i+1])) // add to end
            fnpart(i+1, end, s, prev)
            prev = prev[:len(prev)-1] // remove from end
        }
    }
}

func isPalin(s []rune, start, end int) bool {
    i := start
    j := end
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}