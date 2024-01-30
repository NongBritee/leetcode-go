package leetcode

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type test struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}
	testGroup := []test{
		{NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4}), NewList([]int{7, 0, 8})},
		{NewList([]int{0}), NewList([]int{0}), NewList([]int{0})},
		{NewList([]int{9, 9, 9, 9, 9, 9, 9}), NewList([]int{9, 9, 9, 9}), NewList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, v := range testGroup {
		got := addTwoNumbers(v.l1, v.l2)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("error at %v %v want:%v got:%v", v.l1, v.l2, v.want, got)
		}
	}
}

func NewList(nums []int) *ListNode {
	dummy := new(ListNode)
	result := dummy
	for _, v := range nums {
		result.Next = &ListNode{
			Val: v,
		}
		result = result.Next
	}
	return dummy.Next
}
