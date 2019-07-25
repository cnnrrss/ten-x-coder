func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m+1)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n+1)
    }
    for j := 0; j < n; j++ {
        // we can skip the first row
        // the col vals are already 0
        for i := 0; i < m; i++ {
            if matrix[i][j] == '1' {
                if i == 0 {
                    dp[i][j]++
                } else {
                    dp[i][j] = dp[i-1][j] + 1   
                }
            } else {
                dp[i][j] = 0
            }
        }
    }

    fmt.Println(dp)

    max := 0
    for i := 0; i < m; i++ {
        temp := largestRectangleArea(dp[i], i)
        if max < temp {
            max = temp
        }
    }
    return max
}

func largestRectangleArea(heights []int, i int) int {
    heights = append(heights, -1)
    var maxArea, height, left, right, area int
    var stack []int
    for right < len(heights) {
        if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[right] {
            stack = append(stack, right)
            right++
            continue
        }

        // pop from stack
        height = heights[stack[len(stack)-1]]
        stack = stack[:len(stack)-1]
        //
        if len(stack) == 0 {
            left = -1
        } else {
            left = stack[len(stack)-1]
        }
        area = (right - left - 1) * height

        if maxArea < area {
            maxArea = area
        }
    }
    return maxArea
}