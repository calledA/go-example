package code

import "gmr/gmr-x/ds"

/*
输入一个链表，按链表从尾到头的顺序返回一个ArrayList
*/

//快慢指针
func Num03MethodOne(node *ds.LinkNode) *ds.LinkNode {
	if node == nil {
		return nil
	}
	var slow *ds.LinkNode
	fast := node
	for node != nil {
		fast = node.NextNode
		node.NextNode = slow
		slow = node
		node = fast
	}
	return slow
}

//递归
func Num03MethodTwo(node *ds.LinkNode) *ds.LinkNode {
	if node == nil {
		return nil
	}
	newNode := Num03MethodTwo(node.NextNode)
	node.NextNode.NextNode = node
	node.NextNode = nil
	return newNode
}

//使用栈
func reverseList(head *ds.LinkNode) *ds.LinkNode {
	if head == nil {
		return nil
	}
	myStack := []*ds.LinkNode{}
	for cur := head; cur != nil; {
		temp := cur
		cur = cur.NextNode
		temp.NextNode = nil // 将每个节点的next指针指向空，防止后面死循环
		myStack = append(myStack, temp)

	}
	root := &ds.LinkNode{}
	cur := root // 哨兵节点记录链表的头
	for len(myStack) > 0 {
		root.NextNode = myStack[len(myStack)-1]
		myStack = myStack[:len(myStack)-1]
		root = root.NextNode
	}
	return cur.NextNode
}
