// bfs using 2 queues (1 for x axis, 1 for y axis)

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    nRows, nCols := len(grid), len(grid[0])
    nIslands := 0
    
    xQueue := make([]int, 0, nRows*nCols)
    yQueue := make([]int, 0, nRows*nCols)
    for r := 0; r < nRows; r++ {
        for c := 0; c < nCols; c++ {
            if grid[r][c] == '1' {                
                xQueue = append(xQueue, r)
                yQueue = append(yQueue, c)
                grid[r][c] = '0' // mark as visited
                
                for len(xQueue) > 0 {
                    row, col := xQueue[0], yQueue[0] // dequeue
                    xQueue, yQueue = xQueue[1:], yQueue[1:] // dequeue
                    
                    if 0 <= row-1 && grid[row-1][col] == '1' {
                        xQueue = append(xQueue, row-1)
                        yQueue = append(yQueue, col)
                        grid[row-1][col] = '0'
                    }
                    if row+1 < nRows && grid[row+1][col] == '1' {
                        xQueue = append(xQueue, row+1)
                        yQueue = append(yQueue, col)
                        grid[row+1][col] = '0'
                    }
                    if 0 <= col-1 && grid[row][col-1] == '1' {
                        xQueue = append(xQueue, row)
                        yQueue = append(yQueue, col-1)
                        grid[row][col-1] = '0'
                    }
                    if col+1 < nCols && grid[row][col+1] == '1' {
                        xQueue = append(xQueue, row)
                        yQueue = append(yQueue, col+1)
                        grid[row][col+1] = '0'
                    }                    
                }
                nIslands += 1
            }
        }
    }
    return nIslands
}