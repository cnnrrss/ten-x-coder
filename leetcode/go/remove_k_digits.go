func removeKdigits(num string, k int) string {
    top, digits := 0, len(num) - k
    stack := make([]byte, len(num))
    for i, _ := range num {
        // remove last int of first ascending seq 12[3]21         
        for top > 0 && stack[top-1] > num[i] && k > 0 {
            top--
            k--
        }
        stack[top] = num[i]
        top++
    }
    
    i := 0
    // case of leading zeros in stack i.e. 000123     
    for i < digits && stack[i] == '0' {
        i++
    }
    // case for all zeros
    if i == digits {
        return "0"
    }
    
    return string(stack[i:digits])
}