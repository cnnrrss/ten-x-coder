func sortedSquares(A []int) []int {
    res := make([]int, len(A))
    i, j := 0, len(A)-1
    
    for k := len(A)-1; k >= 0; k-- {
        r1, r2 := A[i] * A[i], A[j] * A[j]
        
        if r1 >= r2 {
            res[k] = r1
            i++
        } else {
            res[k] = r2
            j--
        }
    }
    
    return res
}