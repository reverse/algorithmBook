def Cycle(head): 
    if head is None or head.Next is None:
        return head 

    slow, fast = head, head

    while fast != None and fast.Next != None:
        slow, fast = slow.next, fast.next.next
        if slow == fast:
            return True 

    return False 
