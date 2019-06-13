// |                        |
// |                   _    |
// |       _          | |   |
// |  _   | |    _    | |   |
// |_| |__| |___| |___| |___|

func trap(height []int) int {
    left, right := 0, len(height)-1
    ans := 0
    leftHeight, rightHeight := 0, 0
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftHeight {
                leftHeight = height[left]
            } else {
                ans += leftHeight - height[left]
            }
            left++
        } else {
            if height[right] >= rightHeight {
                rightHeight = height[right]
            } else {
                ans += rightHeight - height[right]
            }
            right--
        }
    }
    return ans
}