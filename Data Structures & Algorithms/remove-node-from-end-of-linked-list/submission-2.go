/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
	current := head

	for current != nil {
		current = current.Next
		length++
	}

	remaining := length - n
	var prev *ListNode
	current = head
	for remaining > 0 {
		if remaining == 1 {
			prev = current
		}
		current = current.Next
		remaining--
	}

	if prev == nil {
		temp := current.Next
		current.Next = nil
		return temp
	}
	
	prev.Next = current.Next
	return head
}
