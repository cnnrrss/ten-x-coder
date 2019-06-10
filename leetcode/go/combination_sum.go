func combinationSum(candidates []int, target int) [][]int {
    ans := [][]int{}
    
    if len(candidates) == 0 || candidates == nil {
        return ans
    }
    
    sort.Ints(candidates)
    set := []int{}
    dfs(candidates, target, 0, &ans, set)
    return ans
}

func dfs(candidates []int, target int, pos int, ans *[][]int, set[]int) {
    if target < 0 {
        return
    }
    if target == 0 {
        c := make([]int, len(set))
        copy(c, set)
        *ans = append(*ans ,c)
        return
    }
    for i := pos; i < len(candidates); i++ {
        set = append(set, candidates[i])
        dfs(candidates, target - candidates[i], i, ans, set)
        set = set[:len(set)-1]
    }
}
