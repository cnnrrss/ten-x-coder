import sys

class Solution:
    def reverse_integer(self, num: int) int ->:
        ans = sys.minint

        while num > 0:
            x , y = divmod(num, 10)
            print(x, y)
