// Brutal Force
func rotate(nums []int, k int)  {
    // k should be remainder of k/length
    temp, orig := 0, 0
    
    k %= len(nums)
    for i := 0; i < k; i++ {
        orig = nums[len(nums)-1]
        for j := 0; j < len(nums); j++ {
            temp = nums[j]
            nums[j] = orig
            orig = temp
        }
    }   
}