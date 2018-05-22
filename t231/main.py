class Solution:
    def isPowerOfTwo(self, n):
        """
        :type n: int
        :rtype: bool
        """
        if n <= 0:
            return False
        while n > 1:
            if n % 2 == 1:
                return False
            n = n / 2
        return True

s = Solution()
print(s.isPowerOfTwo(2))
print(s.isPowerOfTwo(1))
print(s.isPowerOfTwo(0))
print(s.isPowerOfTwo(13))
print(s.isPowerOfTwo(281474976710656))
print(s.isPowerOfTwo(281474976710655))