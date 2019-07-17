func moveZeroes(nums []int)  {
    if len(nums) <= 1 {
        return
    }
    
    lastNonZero := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
            lastNonZero++
        }
    }
    return
}