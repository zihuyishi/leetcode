class Solution:
    def calcSeqs(self, s):
        """
        :type s: str
        :rtype: Dict[str, int]
        """
        seqs = dict()
        for c in s:
            if c.isalpha() == False:
                continue
            c = c.lower()
            if c in seqs:
                seqs[c] += 1
            else:
                seqs[c] = 1
        return seqs
    
    def checkComplete(self, src, dest):
        keys = dest.keys()
        for key in keys:
            if key in src:
                if src[key] < dest[key]:
                    return False
            else:
                return False
        return True

    def shortestCompletingWord(self, licensePlate, words):
        """
        :type licensePlate: str
        :type words: List[str]
        :rtype: str
        """
        seqs = self.calcSeqs(licensePlate)
        small = None
        for word in words:
            wordSeqs = self.calcSeqs(word)
            if self.checkComplete(wordSeqs, seqs):
                if small == None:
                    small = word
                else:
                    if len(word) < len(small):
                        small = word
        return small

s = Solution()
print(s.shortestCompletingWord("13 456", ["looks", "pest", "sew", "show"]))