class BST:
    def __init__(self, value=0, left=None, right=None):
         self.value = value
         self.left = left
         self.right = right

    
def BranchSums(tree):
    sums = []
    branchSumHelper(tree, 0, sums)
    return sums


def branchSumHelper(tree, sum, sums):
    if tree is None:
        return 

    newSum = sum + tree.value

    if tree.left is None and tree.right is None:
        sums.append(newSum)
        return
    
    branchSumHelper(tree.left, newSum, sums)
    branchSumHelper(tree.right, newSum, sums)
