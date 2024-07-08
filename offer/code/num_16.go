package code

import (
	"gmr/gmr-x/ds"
)

/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*/

func Num16MethodOne(l1, l2 *ds.LinkNode) *ds.LinkNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var root = &ds.LinkNode{}
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Data < l2.Data {
				root.NextNode = l1
				l1 = l1.NextNode
			} else if l1.Data > l2.Data {
				root.NextNode = l2
				l2 = l2.NextNode
			} else {
				root.NextNode = l2
				l2 = l2.NextNode
				root.NextNode = l1
				l1 = l1.NextNode
			}
		} else if l1 == nil {
			root.NextNode = l2
			break
		} else if l2 == nil  {
			root.NextNode = l1
			break
		}
		root = root.NextNode
	}
	return root.NextNode
}
