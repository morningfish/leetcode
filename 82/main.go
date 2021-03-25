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
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	PLinkedList(&a)
	x := deleteDuplicates(&a)
	PLinkedList(x)
}
func PLinkedList(h *ListNode) {
	c := h
	for {
		fmt.Println(c.Val)
		if c.Next == nil {
			break
		}
		c = c.Next
	}

}

// 遇到需要操作本位的，向前加一位，永远操作下一位
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}
	// 前面补一位
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		// 从下一位计算，下一位和下下一位相同的话
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			// 本位的下一位去重。
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
