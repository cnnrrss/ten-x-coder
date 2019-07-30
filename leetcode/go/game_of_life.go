func gameOfLife(board [][]int)  {
    // live = 1
    // dead = 0
    // 8 neighbors (horiz, vert, diag)
    // 1. cell = 1 && numAlive(neighbors) < 2; cell = dead (underpop)
    // 2. cell = 1 && numAlive(neighbors) == 2 || 3; cell = live
    // 3. cell = 1 && numAlive(neighbors) > 3; cell = dead (overpop)
    // 4. cell = 0 && numAlive(neighbors) == 3; cell = alive (repro)
    
    //  (11) 3 was alive but now dead
    //  (10) 2 was dead but now alive 
    
    rows := len(board)
    cols := len(board[0])
    
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            alive := livingNeighbors(board, rows, cols, row, col)
            if board[row][col] == 1 && alive >= 2 && alive <= 3 {
                board[row][col] = 3
            }
            if board[row][col] == 0 && (alive == 3) {
                board[row][col] = 2
            }
        }
    }
    
    for m := 0; m < rows; m++ {
        for n := 0; n < cols; n++ {
            board[m][n] >>= 1
        }
    }
}

func livingNeighbors(board [][]int, m, n, i, j int) int {
    alive := 0
    for x := max(i-1, 0); x <= min(i+1,m-1); x++ {
        for y:= max(j-1, 0); y <= min(j+1, n-1); y++ {
            alive += board[x][y] & 1
        }
    }
    alive -= board[i][j] & 1
    return alive
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}