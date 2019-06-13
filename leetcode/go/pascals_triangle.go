func generate(numRows int) [][]int {
    if numRows < 1 {
        return [][]int{}
    }
    
    triangle := make([][]int, numRows)
    triangle[0] = []int{1}
    
    for i:=1; i<numRows; i++ {
        temp := []int{1}
        for j := 1; j < i; j++ {
            temp = append(temp, triangle[i-1][j] + triangle[i-1][j-1])                   
        }
        temp = append(temp, 1)
        triangle[i] = temp
    }
    return triangle
}

//   1
//  1,1
// 1,2,1