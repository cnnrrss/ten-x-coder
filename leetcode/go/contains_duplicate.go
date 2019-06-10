func containsDuplicate(nums []int) bool {
    container := make(map[int]bool)
    for _, n := range nums {
        if _, ok := container[n]; ok {
            return true
        }
        container[n] = true
    }
    return false
}