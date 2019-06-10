# python3

class Solution:
    def findNumberOfLIS(self, nums: List[int]) -> int:
        N = len(nums)         
        if N <= 1: return N # base case
        lengths = [0] * N # lengths[i] = longest ending in nums[i]
        counts = [1] * N # counts[i] = number of longest ending in nums[i]
        
        for j, num in enumerate(nums): # enumerate has auto counter, like range in go
            for i in range(j): # xrange = range in python 3
                if nums [i] < nums[j]:
                    if lengths[i] >= lengths[j]:
                        lengths[j] = 1 + lengths[i]
                        counts[j] =  counts[i]
                    elif lengths[i] + 1 == lengths[j]:
                        counts[j] += counts[i]

        longest = max(lengths)
        sum = 0
        for i, c in enumerate(counts):
            if lengths[i] == longest:
                sum += c
        return sum
