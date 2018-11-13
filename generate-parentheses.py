class Solution1:
    def generateParenthesis(self, n):
        result = []
        def backtrack(s = '', left = 0, right = 0):
            if len(s) == 2 * n:
                result.append(s)
                return
            if left < n:
                backtrack(s + "(", left + 1, right)
            if right < left:
                backtrack(s + ")", left, right + 1)

        backtrack('', 0, 0)
        return result

class Solution2:
    def generateParenthesis(self, n):
        self.result = []
        self.n = n
        self._gen(0, 0, "")
        return self.result

    def _gen(self, left, right, result):
        if left == right == self.n:
            self.result.append(result)
            return
        if left < self.n:
            self._gen(left+1, right, result + "(")
        if right < left:
            self._gen(left, right+1, result + ")")
