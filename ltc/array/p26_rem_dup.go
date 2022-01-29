package array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// [1,1,2]
// [0,0,1,1,1,2,2,3,3,4]
func removeDuplicates(nums []int) int {
   n := len(nums)
   if n == 0 {
	   return 0
   }
   
   j := 1
   for i:=1; i < n ; i++ {
     if nums[i] != nums[i-1] {
		nums[j] = nums[i]
		j++
	 }
   }

   return j
}