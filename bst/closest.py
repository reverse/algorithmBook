class BST:
    def __init__(self, value=0, left=None, right=None):
         self.value = value
         self.left = left
         self.right = right
    

def FindClosestValueInBstRecursive(tree, target):
	return recursiveHelper(tree, target, float("inf"))
	
def recursiveHelper(tree, target, closest):
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

def FindClosestValueInBstIterative(tree, target):
    return iterativeHelper(tree, target, float("inf"))    

def iterativeHelper(tree, target, closest):
    node = tree
    while node is not None:
        if abs(target-closest) > abs(target-node.value):
            closest = node.value
        
        if target < node.value: 
            node = node.left
        elif target > node.value:
            node = node.right
        else:
            break
        
    return closest 
