/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length int
	curr := head 
	for curr != nil {
		length++
		curr = curr.Next
	}

	length -= n

	if length == 0 {
		return head.Next
	}

	curr = head
	for length > 1 {
		curr = curr.Next
		length--
	}

	curr.Next = curr.Next.Next

	return head
}
