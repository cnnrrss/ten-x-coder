func maxSubArray(nums []int) int {
    if len(nums) <= 1 {
        return nums[0]
    }
    
    max := -1<<31
    subTotal := 0
    length := len(nums)
    
    for i:=0; i<length; i++ {
        subTotal += nums[i]
        
        if subTotal < nums[i] {
            subTotal = nums[i]
        }
        
        if max < subTotal {
            max = subTotal
        }
    }
    
    return max
}