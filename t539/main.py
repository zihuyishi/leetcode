class Solution:
    def timeToNumber(self, timePoint):
        items = timePoint.split(':')
        return int(items[0]) * 60 + int(items[1])

    def findMinDifference(self, timePoints):
        """
        :type timePoints: List[str]
        :rtype: int
        """
        smaller = 24 * 60
        timeNums = list(map(self.timeToNumber , timePoints))
        timeNums.sort()
        pre = timeNums[0]
        timeNums.append(pre + 24*60)
        for v in timeNums[1:]:
            if v - pre < smaller:
                smaller = v - pre
            pre = v
        return smaller

s = Solution()
times = ["05:31","22:08","00:35"]
print(s.findMinDifference(times))
