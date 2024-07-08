package code

import "gmr/gmr-x/ds"

//从上往下打印出二叉树的每个节点，同层节点从左至右打印。
func Num22MethodOne(root *ds.BTree) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*ds.BTree{root}
	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		res = append(res, queue[0].Data)
		queue = queue[1:]
	}
	return res
}
