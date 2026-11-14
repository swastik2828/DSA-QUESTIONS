func topKFrequent(nums []int, k int) []int {
      freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    // Step 2: Create buckets
    buckets := make([][]int, len(nums)+1)

    // Step 3: Place numbers into buckets
    for num, count := range freq {
        buckets[count] = append(buckets[count], num)
    }

    // Step 4: Collect top k elements
    result := make([]int, 0, k)

    for i := len(buckets) - 1; i >= 0; i-- {
        for _, num := range buckets[i] {
            result = append(result, num)

            if len(result) == k {
                return result
            }
        }
    }

    return result
}