package linkedList 

type LinkedList struct {
    Val int 
    Next *LinkedList 
}

// Reverse reverses a linked list in place, returns head of reversed LL
func Reverse(head *LinkedList) *LinkedList {
    current := head 
    var prev *LinkedList 

    for current != nil {
        next := current.Next // init our next variable so we can store it while we switch
        current.Next = prev // reverse the next pointer to point to our previous (reversing)
        prev = current // move our previous to what we just switched
        current = next // move forward, give us a new node in the LL
    }
    return prev 
}
