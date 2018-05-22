from collections import deque
class Solution:
    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        if len(image) == 0 or len(image[0]) == 0:
            return image
        queue = deque([[sr,sc]])
        oldColor = image[sr][sc]
        if oldColor == newColor:
            return image
        image[sr][sc] = newColor
        while len(queue) != 0:
            cur = queue.popleft()
            cr, cc = cur[0], cur[1]
            for d in [[1,0],[0,1],[-1,0],[0,-1]]:
                nr = cr + d[0]
                nc = cc + d[1]
                if nr >= 0 and nc >= 0 and nr < len(image) and nc < len(image[0]) and image[nr][nc] == oldColor:
                    image[nr][nc] = newColor
                    queue.append([nr, nc])
        return image

s = Solution()
image = [[0,0,0],[0,1,1]]
print(s.floodFill(image, 1, 1, 1))