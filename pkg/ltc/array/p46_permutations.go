package array

var ans [][]int

func permute(nums []int) [][]int {
    // 1 2 3
    // 1 3 2
    ans = make([][]int, 0)
    permuteR(nums, 0)
    return ans
}

func permuteR(nums []int, idx int) {
    if idx == len(nums) {
        cp := make([]int, len(nums))
        copy(cp, nums)
        ans = append(ans, cp)
        return
    }

// start with idx so that for a string x1x2x3x4 , we do
// x1 -> x2x3x4 in first loop iteration
// x2 -> x2x3x4 in second loop iteration
// x3 -> x2x1x4 
// x4 -> x2x3x1
    for i := idx; i < len(nums); i++ {
        swap(nums, i, idx)
        permuteR(nums, idx + 1)
        swap(nums, i, idx)
    }
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}


// var ans [][]int

// technique with pick/not pick tracked using a map.
func permute1(nums []int) [][]int {
    ans = make([][]int, 0)
    freq := make(map[int]int)
    curr := make([]int, 0)
    permuteR1(nums, curr, freq)
    return ans
}



func permuteR1(nums, curr []int, freq map[int]int) {
    if len(curr) == len(nums) {
        cp := make([]int, len(curr))
        copy(cp, curr)
        ans = append(ans, cp)
        return
    }

    for i := 0; i < len(nums); i++ {
        if freq[i] == 0 {
            freq[i] = 1
            curr = append(curr, nums[i])
            permuteR1(nums, curr, freq)

            // reset
            freq[i] = 0
            curr = curr[:len(curr)-1]
        }
    }
}