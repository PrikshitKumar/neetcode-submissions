/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	
	var result *ListNode
	var curr = result

	for list1 != nil && list2 != nil {
		var temp *ListNode
		if list1.Val <= list2.Val {
			temp = list1
			list1 = list1.Next
		} else {
			temp = list2
			list2 = list2.Next
		}

		if result == nil {
			result = temp
			curr = temp
		} else {
			curr.Next = temp
			curr = curr.Next

		}
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return result
    
}
