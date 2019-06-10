func characterReplacement(s string, k int) int {
    var size, left int
    dict := [128]int{}
    
    for right := 0; right < len(s); right++ {
        dict[s[right]]++
        size = max(size, dict[s[right]])
        
        if right-left+1-size > k {
            dict[s[left]]--
            left++
        }
    }
    
    return len(s) - left
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}