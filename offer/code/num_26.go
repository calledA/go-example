package code

import "gmr/gmr-x/ds"

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
*/

func Num26MethodOne(root *ds.BTree) *ds.BTree {
	if root == nil {
		return nil
	}
	head, tail := dfsLDR(root)
	//最后将返回的头尾节点拼接成环
	tail.Right = head
	head.Left = tail
	return head
}

func dfsLDR(node *ds.BTree) (head *ds.BTree, tail *ds.BTree) {
	if node == nil {
		return
	}
	//递归,左子树
	lHead, lTail := dfsLDR(node.Left)
	if lHead != nil {
		//如果左子树不为空,头结点为左子树的头节点.并拼接当前节点到左子树的尾节点
		head = lHead
		lTail.Right = node
		node.Left = lTail
	} else {
		//左子树为空,头结点为当前节点
		head = node
	}
	//递归,右子树
	rHead, rTail := dfsLDR(node.Right)
	if rTail != nil {
		//如果右子树不为空,尾节点为右子树的尾节点.并拼接当前节点到右子树的头结点
		tail = rTail
		node.Right = rHead
		rHead.Left = node
	} else {
		//右子树为空,尾节点为当前节点
		tail = node
	}
	return
}
