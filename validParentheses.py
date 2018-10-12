class Solution(object):
    def isValid(self, s):
        if len(s) == 0:
            return True
        if len(s) % 2 == 1:
            return False

        charDict = {")":"(", "]":"[", "}":"{"}
        stack = []
        for c in s:
            if c in charDict:
                if not stack or stack.pop() != c:
                    return False
                else:
                    stack.append(c)

        return len(stack) == 0
