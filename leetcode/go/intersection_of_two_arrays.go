func intersection(nums1 []int, nums2 []int) []int {
    var res []int
    idx := make(map[int]bool)
    for _, i := range nums1 {
        idx[i] = false
    }
    for _, i := range nums2 {
        v, ok := idx[i]
        if ok && v == false {
            idx[i] = true
            res = append(res, i)
        }
    }
    return res
}