package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	fmt.Printf("%+v\n", list)
	fmt.Println("11111")
	res := removeElements(&list, 6)
	fmt.Printf("%+v\n", res)
}
func removeElements(head *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Next = head
	for node := tmp; node != nil && node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return tmp.Next
}
