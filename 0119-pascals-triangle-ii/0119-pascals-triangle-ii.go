func getRow(rowIndex int) []int {
    ans := make([]int, 1)
    ans[0] = 1

    for i := 1; i <= rowIndex; i++ {
        ans = append(ans, 1)

        for j := i-1; j > 0; j-- {
            ans[j] = ans[j] + ans[j-1]
        }

    }
    return ans

}