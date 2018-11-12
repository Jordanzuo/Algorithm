# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# This method uses DFS. It's counter intuitive, but is easy to implement.
class Solution1:
    def minDepth(self, root):
        if not root:return 0

        self.min = 100000000
        self.dfs(root, 1)

        return self.min

    def isLeaf(self, node):
        if node.left or node.right:
            return False
        return True

    def dfs(self, root, level):
        if not root:return
        if self.isLeaf(root):
            self.min = level if level < self.min else self.min
        else:
            self.dfs(root.left, level+1)
            self.dfs(root.right, level+1)

# This method uses BFS. It's intuitive, but is not easy to implement.
class Solution:
    def isLeaf(self, node):
        if node.left or node.right:
            return False
        return True

    def minDepth(self, root):
        if not root:return 0

        from collections import deque

        level, isFirstLeaf = 1, False
        queue = deque()
        queue.append(root)

        while queue:
            count = len(queue)
            for _ in range(count):
                item = queue.popleft()
                if self.isLeaf(item):
                    if not isFirstLeaf:
                        return level
                else:
                    if item.left:
                        queue.append(item.left)
                    if item.right:
                        queue.append(item.right)

            level += 1

        return level
