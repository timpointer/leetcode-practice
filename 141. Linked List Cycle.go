package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowP := head
	fastP := head
	for fastP != nil && fastP.Next != nil {
		fastP = fastP.Next.Next
		slowP = slowP.Next
		if slowP == fastP {
			return true
		}
	}
	return false
}

// mark
const mark = 20001

func hasCycle2(head *ListNode) bool {
	for head != nil {
		if head.Val == mark {
			return true
		}
		head.Val = mark
		head = head.Next
	}

	return false
}
