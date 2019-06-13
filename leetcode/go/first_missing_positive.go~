func firstMissingPositive(nums []int) int {
    length, ptr := len(nums), 0
    for ptr < length {
        for nums[ptr] <= length && nums[ptr] > 0 && nums[ptr] != ptr+1 && nums[nums[ptr]-1] != nums[ptr] {
            idx1 := ptr
            idx2 := nums[ptr] - 1
            nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
        }
        ptr++
    }
    for i, num := range nums {
        if num != i+1 {
            return i+1
        }
    }
    return length + 1
}