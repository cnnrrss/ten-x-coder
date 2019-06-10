# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        start, end, pivot, ans = 1, n, 0, 0
        while start <= end:
            # // floor division that results in whole number     
            pivot = (start+end) // 2
            if isBadVersion(pivot):
                ans = pivot
                end = pivot - 1
            else:
                start = pivot + 1
        return ans
