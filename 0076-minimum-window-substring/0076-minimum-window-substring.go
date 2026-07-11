func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }

    // Track character frequencies needed from t
    hash := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        hash[t[i]]++
    }

    // count tracks how many characters from t are fully matched in the window
    count := 0
    minLen := len(s) + 1
    sInd := -1
    l := 0

    for r := 0; r < len(s); r++ {
        // If this character is still needed, it contributes to matching t
        if hash[s[r]] > 0 {
            count++
        }
        // Reduce the requirement for this character
        hash[s[r]]--

        // When all characters from t are matched, try to shrink the window
        for count == len(t) {
            // Fix: Check window length using r-l+1 (not l-r+1)
            if (r - l + 1) < minLen {
                minLen = r - l + 1
                sInd = l
            }

            // Move the left pointer out of the window
            hash[s[l]]++
            // If the character is now strictly required (> 0), we broke the match
            if hash[s[l]] > 0 {
                count--
            }
            l++
        }
    }

    if sInd == -1 {
        return ""
    }
    return s[sInd : sInd+minLen]
}
