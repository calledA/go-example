package code

import "gmr/gmr-x/ds"

/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
*/

func Num24MethodOne(root *ds.BTree, target int) [][]int {
	var path []int
	var res [][]int
	recur(root, target, path, &res)
	return res
}

func recur(root *ds.BTree, target int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	path = append(path, root.Data)
	target = target - root.Data

	if target == 0 && root.Right == nil && root.Left == nil {
		var temp = make([]int, len(path))
		copy(temp, path)
		*res = append(*res, path)
	}
	recur(root.Left, target, path, res)
	recur(root.Right, target, path, res)
}
