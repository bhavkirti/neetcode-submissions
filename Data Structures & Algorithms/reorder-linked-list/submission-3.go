/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    s, f := head, head.Next

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	second := s.Next
	s.Next = nil
	var prev *ListNode

	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	first := head
	second = prev

	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	}
}
