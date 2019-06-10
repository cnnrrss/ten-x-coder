// Bit manipulation solution using XOR
func singleNumber(nums []int) int {
    var result int 
    for _, n := range nums {
        result ^= n
    }
    return result
}

// Map Solution; Uses O(n) space
func singleNumber(nums []int) int {
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    for k, v := range m {
        if v == 1 {
            return k
        }
    }
    return -1
}