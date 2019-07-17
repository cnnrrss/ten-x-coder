func validPalindrome(s string) bool {
    lo, hi := 0, len(s)-1
    return helper([]byte(s), lo, hi, false)
}

func helper(s []byte, lo int, hi int, hasTriedDelete bool) bool {
    for lo < hi {
        if s[lo] != s[hi] {
            if hasTriedDelete {
                return false
            }
            return helper(s, lo+1, hi, true) || helper(s, lo, hi-1, true)
        }
        lo++
        hi--
    }
    return true
}

