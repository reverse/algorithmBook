class BST:
    def __init__(self, val=0, left=None, right=None):
         self.val = val
         self.left = left
         self.right = right
    

def findClosestValueInBst(tree, target):
	return helper(tree, target, float("inf"))
	
def helper(tree, target, closest):
	if tree is None: 
		return closest 
	if abs(target-closest) > abs(target-tree.value):
		closest = tree.value 
	if target < tree.value:
		return helper(tree.left, target, closest)
	elif target > tree.value:
		return helper(tree.right, target, closest)
	else:
		return closest 

