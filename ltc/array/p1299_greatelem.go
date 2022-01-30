package array

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
    n := len(arr)
    
    max := -1
    for i := n-1; i >= 0; i-- {
        tmp := max
        if arr[i] > max {
            max = arr[i]
        }
        
        arr[i] = tmp
    }
    
    return arr
}