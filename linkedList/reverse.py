class LinkedList:
    def __init__(self):
        self.Val = val
        self.Next = next 

def ReverseLinkedListInPlace(head):
    prev, current = None, head

    while current is not None:
        next = current.Next 
        current.Next = prev 
        prev = current
        current = next 
    return prev


def ReverseRecusive(head):
    return helper(head, None)

def helper(head, prev):
    if head is None:
        return prev 

    next = head.Next
    head.Next = prev
    prev = head

    return helper(next, prev)
