func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := Merge(nums1, nums2)
    piv := len(merged)/2
    if len(merged) % 2 == 0 {
        return (float64(merged[piv-1]) + float64(merged[piv])) / 2  
    } else {
        return float64(merged[piv])
    }
    
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

    size, i, j := len(left)+len(right), 0, 0
    slice := make([]int, size, size)

    for k := 0; k < size; k++ {
        if i > len(left)-1 && j <= len(right)-1 {
            slice[k] = right[j]
            j++
        } else if j > len(right)-1 && i <= len(left)-1 {
            slice[k] = left[i]
            i++
        } else if left[i] < right[j] {
            slice[k] = left[i]
            i++
        } else {
            slice[k] = right[j]
            j++
        }
    }
    return slice
}
