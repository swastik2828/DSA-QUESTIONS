
func isPalindrome(x int) bool {
    if x < 0 || x % 10 == 0 && x != 0 {
        return false
    }
    num := x
    ans := 0
    for x > 0 {
        ans = ans * 10 + x % 10
        x /= 10
    }
    if ans == num {
        return true
    }
    return false
}