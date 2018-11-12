# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# This method uses DFS. This is not intuitive, but it's easy to implement.
class Solution1:
    def levelOrder(self, root):
        if not root: return []

        self.result = []
        self.dfs(root, 0)

        return self.result

    def dfs(self, root, level):
        if not root: return

        if len(self.result) < level + 1:
            self.result.append([])
        self.result[level].append(root.val)
        self.dfs(root.left, level+1)
        self.dfs(root.right, level+1)

# This method used BFS. This is intuitive, but it's not easy to implement.
class Solution2:
    def levelOrder(self, root):
	from collections import deque
        if not root: return []

        result, queue = [], deque()
        queue.append(root)

        while queue:
            result.append([])
            count = len(queue)
            for _ in range(count):
                item = queue.popleft()
                result[-1].append(item.val)
                if item.left:
                    queue.append(item.left)
                if item.right:
                    queue.append(item.right)

        return result
