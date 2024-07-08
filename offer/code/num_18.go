package code

import (
	"gmr/gmr-x/ds"
)

/*
	操作给定的二叉树，将其变换为源二叉树的镜像。
*/

//二叉树深度优先遍历，遍历过程中进行左右子树交换操作
func Num18MethodOne(root *ds.BTree) *ds.BTree {
	p := root
	dfs(p)
	return root
}

func dfs(root *ds.BTree) {
	if root == nil {
		return
	}

	//向下递归遍历，过程中交换
	root.Left, root.Right = root.Right, root.Left

	dfs(root.Left)
	dfs(root.Right)
}

//二叉树先进行递归遍历，找到叶子节点，回溯的过程中进行左右子树交换操作。
func Num18MethodTwo(root *ds.BTree) *ds.BTree {
	if root == nil {
		return root
	}

	// 先递归遍历，找到最下面的叶子节点
	left := Num18MethodTwo(root.Left)
	right := Num18MethodTwo(root.Right)

	// 回溯的过程中，从下向上进行交换赋值
	root.Left = right
	root.Right = left
	return root
}
