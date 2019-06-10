func maxArea(height []int) int {    
    high, temp := 0, 0
    for h := 0; h < len(height); h++ {
        for i := len(height)-1; i > h; i-- {
            x := i - h
            temp = x * max(height[h], height[i])
            if temp > high {
                high = temp
            }
        }
    }
    return high
}

func max(x, y int) int {
    if x > y {
        return y
    }
    return x
}