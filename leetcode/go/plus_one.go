func plusOne(digits []int) []int {
    var sum int
    for i := len(digits) -1; i >= 0; i-- {
        sum = digits[i] + 1
        if sum < 10 {
            digits[i] = sum
            return digits
        }
        digits[i] = sum - 10
    }
    return append([]int{1}, digits...)
}