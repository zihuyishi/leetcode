class Solution:

    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        if rowIndex == 0:
            return [1]

        last = self.getRow(rowIndex-1)
        arr = []
        for i in range(rowIndex+1):
            if i == 0 or i == rowIndex:
                arr.append(1)
            else:
                arr.append(last[i-1]+last[i])
        return arr

s = Solution()

ret = s.getRow(7)
print(ret)