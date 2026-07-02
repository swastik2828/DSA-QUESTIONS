func nextPermutation(nums []int)  {
    ind := -1
    n := len(nums)

    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            ind = i
            break
        }
    }

    if ind == -1 {
        slices.Reverse(nums)
        return
    }

    for i := n-1; i > ind; i-- {
        if nums[i] > nums[ind] {
            nums[i], nums[ind] = nums[ind], nums[i]
            break
        }
    }

    slices.Reverse(nums[ind+1:])
}