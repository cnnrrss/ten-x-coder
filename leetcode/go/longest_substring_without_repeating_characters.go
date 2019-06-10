func lengthOfLongestSubstring(s string) int {
    index, res, start, tmp := [128]int{}, 0, 0, 0

    for i, j := range s {
        if start < index[j] {
            start = index[j]
        }

        tmp = i - start + 1

        if res < tmp {
            res = tmp
        }

        index[j] = i + 1
    }
    return res
}
