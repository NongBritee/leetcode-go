package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	result := dummy
	carry := 0
	for {
		sum := carry
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		result.Next = &ListNode{
			Val: sum % 10,
		}
		result = result.Next

		carry = sum / 10

		if l1 == nil && l2 == nil {
			if carry > 0 {
				result.Next = &ListNode{
					Val: carry,
				}
				result = result.Next
			}
			break
		}
	}
	return dummy.Next
}
