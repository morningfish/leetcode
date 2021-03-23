package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	thead := head
	for tmp := n; tmp > 1; tmp -= 1 {
		if thead.Next != nil {
			thead = thead.Next
		} else {
			return head
		}
	}
	// 创建滑动窗口
	var last *ListNode
	this := head
	for thead.Next != nil {
		thead = thead.Next
		last = this
		this = this.Next
	}
	if last == nil {
		return this.Next
	} else {
		last.Next = this.Next
		return head
	}
}
