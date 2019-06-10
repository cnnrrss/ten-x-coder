func reverseString(s []byte)  {
    j := 0
    for i := len(s)-1; i >= len(s)/2; i-- {
        s[i], s[j] = s[j], s[i]
        j++
    }
}