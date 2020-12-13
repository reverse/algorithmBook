def ReverseLinkedListInPlace(head):
    prev, current = None, head

    while current is not None:
        next = current.Next 
        current.Next = prev 
        prev = current
        current = next 

    return prev 
