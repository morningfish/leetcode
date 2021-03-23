package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(t)
	fmt.Println(t.Left)
	fmt.Println(t.Left.Left)
	fmt.Println(t.Left.Right)
	fmt.Println(t.Right)
	fmt.Println(t.Right.Left)
	fmt.Println(t.Right.Right)
	pt := invertTree(&t)
	fmt.Println(pt)
	fmt.Println(pt.Left)
	fmt.Println(pt.Left.Left)
	fmt.Println(pt.Left.Right)
	fmt.Println(pt.Right)
	fmt.Println(pt.Right.Left)
	fmt.Println(pt.Right.Right)

}
func invertTree(root *TreeNode) *TreeNode {
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}
