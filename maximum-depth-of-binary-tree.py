# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# This method uses DFS.It's counter intuitive, but is easy to implement.
class Solution1:
    def maxDepth(self, root):
        if not root: return 0

        self.max = 0
        self.dfs(root, 1)

        return self.max

    def isLeaf(self, node):
        if node.left or node.right:
            return False
        return True

    def dfs(self, root, level):
        if not root: return
        if self.isLeaf(root):
            self.max = level if level > self.max else self.max
        else:
            self.dfs(root.left, level+1)
            self.dfs(root.right, level+1)

# This method uses BFS. It's intuitive, but is not easy to implement.
class Solution:
    def isLeaf(self, node):
        if node.left or node.right:
            return False
        return True

    def maxDepth(self, root):
        from collections import deque

        if not root: return 0

        maxDep, level = 1, 1
        queue = deque()
        queue.append(root)

        while queue:
            count = len(queue)
            for _ in range(count):
                item = queue.popleft()
                if self.isLeaf(item):
                    maxDep = level if level > maxDep else maxDep
                else:
                    if item.left:
                        queue.append(item.left)
                    if item.right:
                        queue.append(item.right)
            level += 1

        return maxDep
