

import "sort"

func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    ans := make([][]int, 0)

    if n < 4 {
        return ans
    }

    sort.Ints(nums)
    t := int64(target)

    for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // Smallest possible sum for this i.
        min := int64(nums[i]) +
            int64(nums[i+1]) +
            int64(nums[i+2]) +
            int64(nums[i+3])

        if min > t {
            break
        }

        // Largest possible sum for this i.
        max := int64(nums[i]) +
            int64(nums[n-3]) +
            int64(nums[n-2]) +
            int64(nums[n-1])

        if max < t {
            continue
        }

        for j := i + 1; j < n-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }

            // Smallest possible sum for this i, j.
            min = int64(nums[i]) +
                int64(nums[j]) +
                int64(nums[j+1]) +
                int64(nums[j+2])

            if min > t {
                break
            }

            // Largest possible sum for this i, j.
            max = int64(nums[i]) +
                int64(nums[j]) +
                int64(nums[n-2]) +
                int64(nums[n-1])

            if max < t {
                continue
            }

            left, right := j+1, n-1

            for left < right {
                sum := int64(nums[i]) +
                    int64(nums[j]) +
                    int64(nums[left]) +
                    int64(nums[right])

                switch {
                case sum < t:
                    left++

                case sum > t:
                    right--

                default:
                    ans = append(ans, []int{
                        nums[i], nums[j],
                        nums[left], nums[right],
                    })

                    left++
                    right--

                    for left < right &&
                        nums[left] == nums[left-1] {
                        left++
                    }

                    for left < right &&
                        nums[right] == nums[right+1] {
                        right--
                    }
                }
            }
        }
    }

    return ans
}
