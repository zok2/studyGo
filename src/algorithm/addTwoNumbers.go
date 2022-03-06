package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) ( head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		}else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main()  {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},

		},
	}
	l2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 8,
			},

		},
	}
	sum := addTwoNumbers(l1,l2)
	for  {
		fmt.Println(sum.Val)
		if sum.Next != nil {
			sum = sum.Next
		}else {
			break
		}
	}
}
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
