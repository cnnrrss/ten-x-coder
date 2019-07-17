func findDuplicates(nums []int) []int {
    numTbl := map[int]int{}
    res := []int{}
    
    for _, n := range nums {
        numTbl[n]++
    }
    
    for i, n := range numTbl {
        if n == 2 {
            res = append(res, i)
        }
    }
    return res
}