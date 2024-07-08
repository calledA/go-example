package ds

import "fmt"

type BTree struct {
	Data  int
	Left  *BTree
	Right *BTree
}

//前序遍历
func DLRTree(root *BTree) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	DLRTree(root.Left)
	DLRTree(root.Right)
}

//中序遍历
func LDRTree(root *BTree) {
	if root == nil {
		return
	}
	LDRTree(root.Left)
	fmt.Println(root.Data)
	LDRTree(root.Right)
}

//后序遍历
func LRDTree(root *BTree) {
	if root == nil {
		return
	}
	LRDTree(root.Left)
	LRDTree(root.Right)
	fmt.Println(root.Data)
}

//层序遍历
// func OrderTree(root *BTree) {
// 	if root == nil {
// 		return
// 	}
// 	queue := new(linkQueue)
// 	queue.Insert(root.Data)

// 	for queue.size > 0 {
// 		element := queue.Remove()
// 		switch ele := element.(type) {
// 		case *BTree:
// 			// 先打印节点值
// 			fmt.Print(ele.Data, "")

// 			// 左子树非空，入队列
// 			if ele.Left != nil {
// 				queue.Insert(ele.Left.Data)
// 			}

// 			// 右子树非空，入队列
// 			if ele.Right != nil {
// 				queue.Insert(ele.Right.Data)
// 			}
// 		}
// 	}
// }

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	Value int
	Times int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (tree *BinarySearchTree) Add(val int) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: val}
		return
	}
	tree.Root.Add(val)
}

func (node *BinarySearchTreeNode) Add(val int) {
	if val < node.Value {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: val}
		} else {
			node.Left.Add(val)
		}
	} else if val > node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: val}
		} else {
			node.Right.Add(val)
		}
	} else {
		node.Times = node.Times + 1
	}
}

func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

func (tree *BinarySearchTree) FindOne(val int) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindOne(val)
}

func (node *BinarySearchTreeNode) FindOne(val int) *BinarySearchTreeNode {
	if node.Value > val {
		if node.Left == nil {
			return nil
		}
		return node.Left.FindOne(val)
	} else if node.Value < val {
		if node.Right == nil {
			return nil
		}
		return node.Right.FindOne(val)
	} else {
		return node
	}
}

func (tree *BinarySearchTree) FindParent(val int) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	if tree.Root.Value == val {
		return nil
	}
	return tree.Root.FindParent(val)
}

func (node *BinarySearchTreeNode) FindParent(val int) *BinarySearchTreeNode {
	if node.Value > val {
		left := node.Left
		if left == nil {
			return nil
		}
		if left.Value == val {
			return node
		} else {
			return left.FindParent(val)
		}
	} else {
		right := node.Right
		if right == nil {
			return nil
		}
		if right.Value == val {
			return node
		} else {
			return right.FindParent(val)
		}
	}
}

/*
删除元素有四种情况：
第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可。
第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除即可。
第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点，这时二叉查找树的性质又满足了。右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到。
第四种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可。
*/
func (tree *BinarySearchTree) Delete(val int) {
	if tree.Root == nil {
		return
	}

	node := tree.Root.FindOne(val)
	if node == nil {
		return
	}

	//查找父节点
	parent := tree.Root.FindParent(val)

	//如果查找的是根节点且没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		//删除的节点有父亲节点但是没有子树

		//没有子树的话，哪边有匹配删除哪边
		if parent.Left != nil && val == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		//如果左右子树都有，用右子树的最小值作为节点
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		tree.Delete(minNode.Value)
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		//只有一颗子树，将该子树替换被删除节点即可
		//父亲为空，删除的是根节点，删除根节点切换子树为根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		//左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && val == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && val == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}
