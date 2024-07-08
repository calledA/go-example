package code

import (
	"gmr/gmr-x/ds"
)

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
*/

func Num04MethodOne(preOrder, midOrder []int) (root *ds.BTree) {
	if len(preOrder) == 0 {
		return nil
	}
	//根据前序遍历的性质在中序遍历中找到根节点的位置
	var i = 0
	for i < len(preOrder) && preOrder[0] != midOrder[i] {
		i++
	}
	l := Num04MethodOne(preOrder[1:i+1], midOrder[:i])
	r := Num04MethodOne(preOrder[i+1:], midOrder[i+1:])
	return &ds.BTree{
		Data:  preOrder[0],
		Left:  l,
		Right: r,
	}
}
