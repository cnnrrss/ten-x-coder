# python3

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        counter = [0] * 26
        for i in s:
            counter[ord(i)-ord('a')] += 1
        for i in t:
            counter[ord(i)-ord('a')] -= 1
        for count in counter:
            if count != 0:
                return False
        return True
