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
