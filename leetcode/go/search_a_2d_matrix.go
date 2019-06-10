func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    
    tRow, tCol := len(matrix), len(matrix[0])
    low, high := 0, tRow * tCol - 1
    
    var row, col int
    for low <= high {
        mid := (low + high) / 2
        row = mid / tCol
        col = mid % tCol
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}

func searchMatrix2(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    col, row := len(matrix[0])-1, 0
    if target < matrix[0][0] {
        return false
    }
    if target > matrix[len(matrix)-1][col] {
        return false
    }
    lastRow := len(matrix) - 1
    for {
        if col < 0 || row > lastRow {
            return false
        } else if target == matrix[row][col] {
            return true
        } else if target > matrix[row][col] {
            row++
        } else if target < matrix[row][col] {
            col--
        }
    }
}