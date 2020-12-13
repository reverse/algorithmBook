package linkedList 

// MiddleNode takes a linkedList and returns the middle node 
func MiddleNode(head *LinkedList) *LinkedList {
    count := 0 
    current := head 

    for current != nil {
        count++
        current = current.Next 
    }
    
    count = int(count/2)
    pointerTwo := 0 

    for head != nil {
        if pointerTwo == count {
            return head 
        }
        head = head.Next 
        count++ 
    }

    return head 
}

func MiddleNodeOneLiner(head *LinkedList) *LinkedList {
    if head == nil || head.Next == nil {
        return head 
    }
    slow, fast := head, head 
    
    for fast != nil && fast.Next != nil { 
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow

}
