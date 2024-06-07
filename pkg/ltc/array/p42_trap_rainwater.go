package array

// start to MAX and MAX to end, track prevSmall and prevBig as we iterate.
// O(n) time, O(1) mem
func trap(height []int) int {

    max := -1
    maxidx := -1
    for i := 0; i < len(height); i++ {
        if height[i] >= max {
            max = height[i]
            maxidx = i
        }
    }

    trapped := 0

    
    // till last max, find next smallest element on left side
    prevSmall := 0
    for i := 1; i <= maxidx; i++ {
        if height[i] < height[prevSmall] {
            continue
        }
        for j := prevSmall + 1; j < i; j++ {
            trapped = trapped + (height[prevSmall] - height[j])
        }
        prevSmall = i
    }

    prevBig := len(height) - 1
    for i := len(height)-2; i >= maxidx; i-- {
        if height[i] < height[prevBig] {
            continue
        }
        for j := prevBig-1; j > i; j-- {
            trapped = trapped + (height[prevBig] - height[j])
        }
        prevBig = i
    }

    return trapped

}
func trap1(height []int) int {
    trapped := 0
    max := -1
    maxidx := -1
    for i := 0; i < len(height); i++ {
        if height[i] >= max {
            max = height[i]
            maxidx = i
        }
    }

// till last max, find biggest element on right side
    for i := 0; i <= maxidx; {
        nextBig := findEqualOrNextBigger(height, i)
        if nextBig == -1 {
            i++
            continue
        }
        // start from current to next big and find trapped water
        for j := i + 1; j < nextBig; j++ {
            trapped = trapped + (height[i] - height[j])
        }
        // start from nextBug
        i = nextBig
    }

    // after last max, find biggest element smaller than current
    for i := maxidx; i < len(height); {
        lowerBound := findLowerBound(height, i)
        // fmt.Printf("idx: %d, lowerBound: %d\n", i, lowerBound)
        if lowerBound == -1 {
            i++
            continue
        }
        for j := i + 1; j < lowerBound; j++ {
            trapped = trapped + (height[lowerBound] - height[j])
        }
        i = lowerBound
    }

    return trapped
}

// return index of equal number or bigger number smaller than current
func findLowerBound(ar []int, curr int) int {
    nextBig := -1
    for i := curr + 1; i < len(ar); i++ {
        if nextBig == -1 {
            nextBig = i
        } else if ar[i] >= ar[nextBig] {
            nextBig = i
        }
    }
    return nextBig
}

// return index of next big number, -1 if it does not exist
func findEqualOrNextBigger(ar []int, curr int) int {
    for i := curr + 1; i < len(ar); i++ {
        if ar[i] >= ar[curr] {
            return i
        }
    }

    return -1;
}