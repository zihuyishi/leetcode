"""
implement a Trie for fun, more on wiki:
https://en.wikipedia.org/wiki/Trie
"""
import functools

class Node:
    def __init__(self, char):
        self.char = char
        self.children = dict()
        self.endpoint = False
    
    def getCharacter(self, char):
        if char in self.children:
            return self.children[char]
        else:
            return None

    def appendCharacter(self, char):
        c = self.getCharacter(char)
        if c == None:
            c = Node(char)
            self.children[char] = c
            return c
        else:
            return c

class Trie:
    def __init__(self):
        self.root = Node('*')
    
    def append(self, word):
        cur = self.root
        for c in word:
            cur = cur.appendCharacter(c)

        cur.endpoint = True

    def remove(self, word):
        cur = self.root
        for c in word:
            cur = cur.getCharacter(c)
            if cur == None:
                return False
        ret = cur.endpoint
        cur.endpoint = False
        return ret

    def exists(self, word):
        cur = self.root
        for c in word:
            cur = cur.getCharacter(c)
            if cur == None:
                return False

        return cur.endpoint

    def words(self):
        ret = []
        stack = []
        stack.append((self.root, ''))
        while len(stack) > 0:
            cur = stack.pop()
            node = cur[0]
            curStr = cur[1]
            if node.endpoint:
                ret.append(curStr)
            for c, v in node.children.items():
                stack.append((v, curStr+c))
        
        return ret

    def gc(self):
        self.__pathHasWord(self.root)

    def space(self):
        """
        :rtype: int
        """
        return self.__space(self.root)
    
    def __space(self, node):
        """
        :type node: Node
        :rtype: int
        """
        counts = [self.__space(n) for _, n in node.children.items()]
        return len(node.children) + functools.reduce(lambda acc, x: acc + x, counts, 0)

    def __pathHasWord(self, node):
        """
        :type node: Node
        :rtype: bool
        """
        has = node.endpoint
        keys = node.children.keys()
        copyKeys = [key for key in keys]
        for key in copyKeys:
            if self.__pathHasWord(node.children[key]):
                has = True
            else:
                del node.children[key]
        return has

if __name__ == '__main__':
    t = Trie()
    t.append('asdf')
    t.append('bababa')
    t.append('hehe')
    t.append('asfd')
    t.remove('asfd')
    print(t.space())
    t.gc()
    print(t.space())

    keys = t.words()
    print(keys)
    e = [t.exists(w) for w in keys]
    print(e)

    t.remove('asdf')
    t.remove('hehe')
    t.remove('bababa')
    t.gc()
    print(t.space())
