func majorityElement(nums []int) []int {
    var ans []int
    mpp := make(map[int]int)

    for i := 0; i< len(nums); i++ {
        mpp[nums[i]]++
    }

    for key, value := range mpp {
        if value > len(nums)/3 {
            ans=append(ans,key)
        }
    }
    return ans
}