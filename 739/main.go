package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow,fast:= head,head
	for fast != nil && fast.Next != nil {         //找中点
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {                          //奇数链表取下一点
		slow = slow.Next
	}
	var pre *ListNode
	for slow != nil {                           //后半段反转链表
		slow, pre,slow.Next = slow.Next,slow,pre
	}
	for pre != nil {                                 //和原来的比较
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true                //总共O(2n）/ O(2)
}


func main() {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	res := isPalindrome(&list)
	fmt.Println(res)
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	fmt.Println( isPalindrome(&list2))
}
