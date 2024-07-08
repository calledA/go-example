package code

import "gmr/gmr-x/ds"

/*
输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点），请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
*/

func Num25MethodOne(root *ds.RandomListNode) *ds.RandomListNode {
	if root == nil {
		return nil
 }

 //从A开始
 currentNode := root
 //1.复制结点A得到A1，A1放到A后面
 for currentNode != nil {

		//先存一下，nextNode现在为B
		nextNode := currentNode.Next
		cloneNode := new(ds.RandomListNode)
		cloneNode.Data = currentNode.Data
		cloneNode.Next = nextNode
		currentNode.Next = cloneNode
		//A处理完，处理B
		currentNode = nextNode
 }

 //从A开始
 currentNode = root
 //2.复制Random，A到C，A1到C1
 for currentNode != nil {
		if currentNode.RandomNext == nil {
			 currentNode.Next.RandomNext = nil
		} else {
			 currentNode.Next.RandomNext = currentNode.RandomNext.Next
		}
		//A处理完，处理B
		currentNode = currentNode.Next.Next
 }

 //从A开始
 currentNode = root
 pCloneHead := root.Next
 //3.拆分，ABC，A1B1C1
 for currentNode != nil {
		//A1
		cloneNode := currentNode.Next
		currentNode.Next = cloneNode.Next
		if cloneNode.Next != nil {
			 cloneNode.Next = cloneNode.Next.Next
		}
		//A处理完，处理B
		currentNode = currentNode.Next
 }
 return pCloneHead
}

