package dp

func rob2(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    // rob first, skip last Or skip first, rob last
    return max(dpmem(nums[1:]), dpmem(nums[:n-1]))
}
// same dp from house_robber, but passed input is different
func dpmem(nums []int) int {
    n := len(nums)
    mem := make([]int, n)

    // reversing the equation of mem from recmem()
    // mem[i] = max(mem[i-2] + nums[i], mem[i-1])
    // mem[<0] = 0
    // mem[0] = max(mem[-2] + nums[0], mem[-1]) = nums[0]
    // mem[1] = max(mem[-1] + nums[1], mem[0]) = nums[0]
    mem[0] = nums[0]
    if n < 2 {
        return mem[n-1]
    }
    mem[1] = max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        mem[i] = max(mem[i-2] + nums[i], mem[i-1])
    }
    return mem[n-1]
}