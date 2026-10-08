func removeDuplicates(nums []int) int {
    slow := 0
    n := len(nums)

    for fast := 1; fast < n; fast ++ {
        if nums[fast] != nums[slow]{
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}