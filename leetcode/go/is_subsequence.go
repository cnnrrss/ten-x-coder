func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    if t == "" {
        return false
    }
    
    index := 0
    for i := range t {
        if index < len(s) {
            if t[i] == s[index] {
                index++
            }
        }
    }
    
    if index == len(s) {
        return true
    }
    return false
}