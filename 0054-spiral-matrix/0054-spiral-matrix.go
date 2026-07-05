func spiralOrder(mat [][]int) []int {
	// Handle empty matrix edge case safely
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}

	n := len(mat)
	m := len(mat[0])
	left, right := 0, m-1
	top, bottom := 0, n-1
	var ans []int

	for top <= bottom && left <= right {
		// Traverse right
		for i := left; i <= right; i++ {
			ans = append(ans, mat[top][i])
		}
		top++

		// Traverse down
		for i := top; i <= bottom; i++ {
			ans = append(ans, mat[i][right])
		}
		right--

		// Traverse left
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, mat[bottom][i])
			}
			bottom--
		}

		// Traverse up
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, mat[i][left])
			}
			left++
		}
	}
	return ans
}
