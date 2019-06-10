func canJump(nums []int) bool {
    toJump := len(nums) - 1
    for i := toJump; i >= 0; i-- {
        if i + nums[i] >= toJump {
            toJump = i
        }
    }
    return toJump == 0
}