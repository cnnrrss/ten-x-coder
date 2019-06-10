# python3

class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        result = []
        carry = 1
        for dig in digits[::-1]:
            result.append( (dig + carry) % 10 )
            carry = (dig + carry) // 10
        if carry:
            result.append(1)
        return result[::-1]
