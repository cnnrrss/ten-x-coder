func permute(nums []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}

    helper = func(nums []int, n int){
        if n == 1 {
            tmp := make([]int, len(nums))
            copy(tmp, nums)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(nums, n-1)
                if n % 2 == 1 {
                    tmp := nums[i]
                    nums[i] = nums[n-1]
                    nums[n - 1] = tmp
                } else {
                    tmp := nums[0]
                    nums[0] = nums[n-1]
                    nums[n-1] = tmp
                }
            }
        }
    }
    
    helper(nums, len(nums))
    return res
}