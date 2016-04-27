import random
from time import time


class BinaryNode(object):
    
    def __init__(self, value=None):
        """
        Create a binary node
        """
        self.value = value
        self.left = None
        self.right = None

    def add(self, value):
        """
        Adds new node to the three containing this value
        """
        if value <= self.value:
            if self.left:
                self.left.add(value)
            else:
                self.left = BinaryNode(value)
        else:
            if self.right:
                self.right.add(value)
            else:
                self.right = BinaryNode(value)

    def delete(self):
        """
        Remove value from binary tree
        """
        if self.left == self.right is None:
            return None
        if self.left is None:
            return self.right
        if self.right is None:
            return self.left

        child = self.left
        grandchild = self.right

        if grandchild:
            while grandchild.right:
                child = grandchild
                grandchild = child.right

            self.value = grandchild.value
            child.right = grandchild.left
        else:
            self.left = child.left
            self.value = child.value

        return self


class BinaryTree(object):

    def __init__(self):
        """
        Create empty binary tree
        """
        self.root = None

    def add(self, value):
        """
        Insert value into proper location in BR
        """
        if self.root is None:
            self.root = BinaryNode(value)
        else:
            self.root.add(value)

    def contains(self, target):
        """
        Check BT contains target value
        :param target: value for search
        :return: True if contain False if not
        """
        node = self.root
        while node:
            if target == node.value:
                return True
            if target < node.value:
                node = node.left
            else:
                node = node.right
        else:
            return False

    def remove(self, value):
        """
        Remove value from tree
        """
        if self.root:
            self.root = self.remove_from_parent(self.root, value)

    def remove_from_parent(self, parent, value):
        """
        remove tree from tree rooted at parent
        :param parent: parent Node of this value
        :param value: value for removw
        :return: parent or None or delete parent
        """
        if parent is None:
            return None

        if value == parent.value:
            return parent.delete()
        elif value < parent.value:
            parent.left = self.remove_from_parent(parent.left, value)
        else:
            parent.right = self.remove_from_parent(parent.right, value)

        return parent


def performance():
    """
    check performance
    """
    n = 1024
    while n < 9999999:
        bt = BinaryTree()
        for i in range(n):
            bt.add(random.randint(1, n))

        now = time()
        bt.contains(random.randint(1, n))
        print (n, (time() - now)*1000)

        n *= 2

if __name__ == '__main__':
    performance()
