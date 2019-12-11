from __future__ import print_function
import queue

try:
    raw_input = input  # Python 3

class TreeNode:
    def __init__(self, data):
        self.data = data
        self.left, self.right = None, None

def build_tree():
    print("\n********Press N to stop entering at any point of time********\n")
    print("Enter the value of the root node: ", end="")

    check = raw_input().strip().lower()
    if check == 'n':
        return None

    data = int(check)
    q = queue.Queue()
    node = TreeNode(data)
    q.put(node)

    while not q.empty():
        found = q.get()
        print("Enter the left node of %s: " % found.data, end="")

        check = raw_input().strip().lower()
        if check == 'n':
            return node

        left_data = int(check)
        left_node = TreeNode(left_data)
        found.left = left_node
        q.put(left_node)
        print("Enter the right node of %s: " % found.data, end="")

        check = raw_input().strip().lower()

        if check == 'n':
            return node

        right_data = int(check)
        right_node = TreeNode(right_data)
        found.right = right_node
        q.put(right_node)


def pre_order(root):
    # return if invalid node
    if not isinstance(root, TreeNode) or not root:
        return
    print(root.data, end=" ")
    pre_order(root.left)
    pre_order(root.right)


def in_order(root):
    if not isinstance(root, TreeNode) or not root:
        return
    pre_order(root.left)
    print(root.data, end=" ")
    pre_order(root.right)


def post_order(root):
    if not isinstance(root, TreeNode) or not root:
        return

    pre_order(root.left)
    pre_order(root.right)
    print(root.data, end=" ")

def bfs(root):
    if not isinstance(root, TreeNode) or not root:
        return

    q = queue.Queue()
    q.put(root)

    while not q.empty():
        node = q.get()
        print(node.data, end=" ")
        if node.left:
            q.put(node.left)
        if node.right:
            q.put(node.right)

def dfs(root):
    # dfs is a pre order iterative traversal
    if not isinstance(root, TreeNode) or not root:
        return

    stack = []
    node = root

    while node or stack:
        while node:
            print(node.data, end=" ")
            stack.append(node)
            node = node.left
        node = stack.pop()
        node = node.right()
