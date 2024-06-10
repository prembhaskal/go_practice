package dp

func maxProduct(nums []int) int {
    // +ve
    // total max
    tmax := -10000000
    // prev max and min
    pmax := 1
    pmin := 1
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        if curr < 0 {
            pmax, pmin = pmin, pmax
        }

        // either use prev number or start fresh
        // this works for both when curr is 0
        // when curr is +ve and prev max itself was -ve -> [-2, 3]
        // when curr is -ve and prev min was 0 -> [0, -2]
        pmax = max(pmax * curr, curr)
        pmin = min(pmin * curr, curr)
        
        tmax = max(tmax, pmax)
    }

    return tmax
}

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func min(a, b int)int {
//     if a < b {
//         return a
//     }
//     return b
// }