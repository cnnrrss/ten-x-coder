// twoSum solution requires O(n) space
func twoSum(nums []int, target int) []int {
    solMap := map[int]int{}
    
    for i := 0; i < len(nums); i++ {
        comp := target - nums[i]
        if sol, ok := solMap[comp]; ok {
            return []int{sol, i} 
        }
        solMap[nums[i]] = i
    }
    return []int{}
}