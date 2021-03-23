package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	b := swapPairs(&a)

	fmt.Println(b.Val, b.Next.Val, b.Next.Next.Val)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := ListNode{}
	newNode.Val = head.Next.Val
	newNode.Next = &ListNode{
		Val:  head.Val,
		Next: swapPairs(head.Next.Next),
	}
	return &newNode
}
