func peakIndexInMountainArray(A []int) int {
    if len(A) < 3 {
        return 0
    }   
    lo, hi := 0, len(A) - 1
    for lo < hi {
        mid := lo + (hi-lo) / 2
        if A[mid] < A[mid+1] {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    return lo
}

