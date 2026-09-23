/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		kth := getKthElement(groupPrev, k)
		if kth == nil {
			break
		}

		groupNext := kth.Next
		prev, curr := groupNext, groupPrev.Next
		for curr != groupNext {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}
		tmp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = tmp
	}
	return dummy.Next
    
}

func getKthElement(curr *ListNode, k int) *ListNode{
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}
