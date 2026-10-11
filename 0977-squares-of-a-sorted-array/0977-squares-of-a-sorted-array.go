func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    left, right := 0, n-1
    index := n-1
    for left <= right {
        ls := nums[left]*nums[left]
        rs := nums[right]*nums[right]
        
        if ls > rs {
            res[index] = ls
            left++
        } else {
            res[index] = rs
            right--
        }
        index--
    }
    return res
}