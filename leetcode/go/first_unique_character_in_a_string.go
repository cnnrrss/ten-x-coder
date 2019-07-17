func firstUniqChar(s string) int {
    dic := map[rune]int{}
    for _, r := range s {
        dic[r]++
    }
    
    for i, r := range s {
        if dic[r] == 1 {
            return i
        }
    }
    
    return -1
}