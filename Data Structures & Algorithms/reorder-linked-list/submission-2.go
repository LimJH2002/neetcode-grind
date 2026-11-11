/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverseList(slow.Next)
	first := head
	slow.Next = nil

	for first != nil && second != nil {
		temp1 := first.Next
		temp2 := second.Next
		first.Next = second
		second.Next = temp1
		first = temp1
		second = temp2
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new_head := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return new_head
}