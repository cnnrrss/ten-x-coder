func numTrees(n int) int {
    if n == 0 {
        return 1
    } else if n == 1 {
        return 1
    }

    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    
    for level := 2; level <= n; level++ {
        for root := 1; root <= level; root++ {
            dp[level] += dp[level-root] * dp[root-1] 
        }
    }
    
    return dp[n]
}