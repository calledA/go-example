package code

//二叉搜索树的后序遍历序列
//从上往下打印出二叉树的每个节点，同层节点从左至右打印。
/*
根节点的值大于左子树小于右子树
后序遍历: 左子树 -> 右子树 -> 根节点
*/
func Num23MethodOne(postorder []int) bool {
	if postorder == nil || len(postorder) <= 0 {
		return false
	}
	length := len(postorder)
	root := postorder[length-1]
	l := 0
	for ; l < length-1; l++ {
		if root < postorder[l] {
			break
		}
	}
	r := l
	for ; r < length-1; r++ {
		if root > postorder[r] {
			return false
		}
	}
	left := true
	if l > 0 {
		left = Num23MethodOne(postorder[:l])
	}
	right := true
	if r < length-1 {
		right = Num23MethodOne(postorder[l:length])
	}
	return left && right
}
