package code

import "gmr/gmr-x/ds"

/*
输入一个链表，输出该链表中倒数第k个结点。
*/

//利用快慢指针的思想。将第一个指针 fast 指向链表的第 k + 1 个节点，第二个指针 slow 指向链表的第一个节点，此时指针 fast 与 slow 之间刚好间隔 k 个节点。此时两个指针同步向后走，当第一个指针 fast 走到链表的尾部空节点时，则此时 slow 指针刚好指向链表的倒数第 k 个节点。
func Num14MethodOne(head *ds.LinkNode, k int) *ds.LinkNode {
	if head == nil {
		return head
	}
	var slow, fast = head, head
	for k > 0 {
		fast = fast.NextNode
		k--
	}

	for fast != nil {
		fast = fast.NextNode
		slow = slow.NextNode
	}
	return slow
}
