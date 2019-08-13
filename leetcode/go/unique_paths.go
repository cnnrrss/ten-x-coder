func uniquePaths(m int, n int) int {
    if m == 0 || n == 0 {
        return 0
    }
    
    // Create a 2D table to store results of subproblems 
    path := [100][100]int{}

    // Count of paths to reach any cell in first row is 1 
    for i := 0; i < m; i++ {
        path[i][0] = 1
    }
    
    // Count of paths to reach any cell in first column is 1 
    for j := 0; j < n; j++ {
        path[0][j] = 1
    }
    
    // Calculate count of paths for other cells in bottom-up
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            path[i][j] = path[i-1][j] + path[i][j-1]
        }       
    }

    return path[m-1][n-1]
}