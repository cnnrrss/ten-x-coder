func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    for i := 1; i < len(nums); i++ {
        fmt.Println(nums[i-1], nums[i], i)
        if nums[i] != nums[i-1] {
            if nums[i] == nums[i+1] {
                return nums[i-1]   
            }
            return nums[i]
        }
        
        i++
        
        if i == len(nums) - 1 {
            return nums[i]
        }
    }
    return 0
}