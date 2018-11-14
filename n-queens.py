#!/usr/bin/python3

class Solution1:
    def solveNQueens(self, n):
        res = []
        self.dfs(0, [], set(), set(), set(), n, res)
        return res

    def dfs(self, row, validCols, usedCols, usedPies, usedNas, n, res):
        if row == n:
            res.append(["."*i+"Q"+"."*(n-i-1) for i in validCols])
            return

        for col in range(n):
            if col not in usedCols and row + col not in usedPies and col - row not in usedNas:
                validCols.append(col)
                usedCols.add(col)
                usedPies.add(row+col)
                usedNas.add(col-row)
                self.dfs(row+1, validCols, usedCols, usedPies, usedNas, n, res)
                validCols.pop(-1)
                usedCols.remove(col)
                usedPies.remove(row+col)
                usedNas.remove(col-row)

class Solution2:
    def solveNQueens(self, n):
        self.n = n
        self.result = list()
        self.dfs(0, [x for x in range(n)] ,self.initTwoDimentionalArray())
        return self.result

    def dfs(self, level, candidates, pos):
        if len(candidates) == 0: 
            self.reset()
            return
        for item in candidates:
            self.choose_kill(pos, level, item)
            if level+1 == self.n:
                self.result.append(self.traverse(pos))
                self.reset()
                continue
            _candidates = self.chooseCandidate(pos, level+1)
            self.dfs(level+1, _candidates, pos)

    def initTwoDimentionalArray(self):
        pos = list()
        for i in range(self.n):
            item = list()
            for j in range(self.n):
                item.append("0")
            pos.append(item)
        return pos

    def reset(self, pos):
        for i in range(self.n):
            for j in range(self.n):
                pos[i][j] = "0"

    def cloneTwoDimentionalArray(self, pos):
        newPos = list()
        for i in range(self.n):
            item = list()
            for j in range(self.n):
                item.append(pos[i][j])
            newPos.append(item)
        return newPos

    def chooseCandidate(self, pos, row):
        candidates = list()
        for col in range(self.n):
            if pos[row][col] == "0":
                candidates.append(col)
        return candidates

    def choose_kill(self, pos, row, col):
        for _row in range(row, self.n):
            for _col in range(self.n):
                if _row == row or _col == col:
                    if _row == row and _col == col:
                        pos[row][col] = "Q"
                    else:
                        pos[_row][_col] = "."
                elif abs((row - _row)/(col - _col)) == 1:
                    pos[_row][_col] = "."

    def traverse(self, pos):
        result = list()
        for row in range(self.n):
            result.append("".join(pos[row]))

        return result

if __name__ == "__main__":
    solution = Solution()
    result = solution.solveNQueens(6)
    print("Final result", result)
