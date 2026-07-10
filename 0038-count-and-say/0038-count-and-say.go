func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }

    result := "1"

    for i := 2; i <= n; i++ {
        var builder strings.Builder
        count := 1

        for j := 1; j < len(result); j++ {
            if result[j] == result[j-1] {
                count++
            } else {
                builder.WriteString(strconv.Itoa(count))
                builder.WriteByte(result[j-1])

                count = 1
            }
        }

        builder.WriteString(strconv.Itoa(count))
        builder.WriteByte(result[len(result) - 1])

        result = builder.String()
    }

    return result
}