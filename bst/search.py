class BST:
    def __init__(self, value=0, left=None, right=None):
         self.value = value
         self.left = left
         self.right = right
    

def SearchBSTIterative(tree, val):
    while tree is not None:
        if val < tree.value:
            tree = tree.left
        elif val > tree.value:
            tree = tree.right
        else:
            return tree
    
    return tree