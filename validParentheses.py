class Solution(object):
    def isValid(self, s):
        length = len(s)
        if length == 0:
            return True
        if length & 1:
            return False

        charDic = {")":"(", "]":"[", "}":"{"}
        stack = []
        for c in s:
            if c in charDic:
                if len(stack) == 0 or stack.pop() != charDic[c]:
                    return False
            else:
                stack.append(c)

        return len(stack) == 0
