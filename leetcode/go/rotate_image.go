func rotate(m [][]int) {
    n := len(m)
    for r := 0; r < n/2; r++ {
        for c := r; c < n-r-1; c++ {            
            m[r][c], m[n-c-1][r], m[n-r-1][n-c-1], m[c][n-r-1] = m[n-c-1][r], m[n-r-1][n-c-1], m[c][n-r-1], m[r][c]
        }
    }
}

