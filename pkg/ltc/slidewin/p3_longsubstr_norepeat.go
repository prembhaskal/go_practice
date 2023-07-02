package slidewin
func lengthOfLongestSubstring(s string) int {
    idxmap := make(map[rune]int) // max idx where ch is seen.
    chrs := []rune(s)
    left := 0
    mxlen := 0
    for idx, ch := range chrs {
        lastidx,ok := idxmap[ch]
        if ok && lastidx >= left {
            left = lastidx + 1
        }
        ln := idx - left + 1
        idxmap[ch] = idx
        mxlen = max3(mxlen, ln)
    
        // fmt.Printf("at idx: %d, ch: %c, mxlen: %d, left: %d\n", idx, ch, mxlen, left)
    }

    return mxlen
}

func lengthOfLongestSubstring2(s string) int {
    seen := make(map[rune]int)
    str := []rune(s)

    start := 0
    if len(s) == 0 {
        return 0
    }
    seen[str[0]] = 0
    end := 1
    mx := 1
    ln := 1
    for ;end < len(str);end++ {
        ch := str[end]

        if _, ok := seen[ch]; ok {
            tmp := seen[ch]
            // we need to remove prev ones too
            for ; start < tmp + 1; start++ { // remove loop,optimize with >= start check on index,
                delete(seen, str[start])
            }
            start = tmp + 1
        }

        seen[ch] = end

        // capture length
        ln = end - start + 1
        mx = max3(mx, ln)
    }
    return mx
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
