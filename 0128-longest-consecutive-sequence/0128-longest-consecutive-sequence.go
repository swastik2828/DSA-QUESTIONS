func longestConsecutive(nums []int) int {
    set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}

	best := 0
	for n := range set {
		if _, ok := set[n-1]; ok {
			continue
		}

		curr := n
		length := 1
		for {
			if _, ok := set[curr+1]; !ok {
				break
			}
			curr++
			length++
		}

		if length > best {
			best = length
		}

	
	}
	return best
}