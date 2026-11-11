/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	if len(lists) < 2 {
		return lists[0]
	}

	first := mergeKLists(lists[:len(lists)/2])
	second := mergeKLists(lists[len(lists)/2:])

	return merge(first, second)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = merge(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge(list1, list2.Next)
		return list2
	}
}
