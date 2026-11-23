func characterReplacement(s string, k int) int {
    var count[26] int
	left, maxFreq := 0, 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		count[idx]++

		if count[idx] > maxFreq {
			maxFreq = count[idx]
		}

		if (right - left+1) - maxFreq > k {
			count[s[left]-'A']--
			left++
		}
 	}
	return len(s) - left
}