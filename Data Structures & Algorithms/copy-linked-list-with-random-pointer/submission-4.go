/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := map[*Node]*Node{nil: nil}

	curr := head

	for curr != nil {
		cop := &Node{Val: curr.Val}
		oldToCopy[curr] = cop
		curr = curr.Next
	}
	curr = head
	for curr!=nil {
		cop := oldToCopy[curr]
		cop.Next = oldToCopy[curr.Next]
		cop.Random = oldToCopy[curr.Random]
		curr = curr.Next
	}

	return oldToCopy[head]
}
