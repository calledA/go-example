package code

/*
完成队列的Push和Pop操作。 队列中的元素为int类型。
*/

type ArrayStack struct {
	Array []string
	Size  int
}

func (stack *ArrayStack) Push(val string) {
	stack.Array = append(stack.Array, val)
	stack.Size = stack.Size + 1
}

func (stack *ArrayStack) Pop() string {
	if stack.Size == 0 {
		panic("stack empty")
	}
	val := stack.Array[0]
	newArray := make([]string, stack.Size-1)
	for i := 0; i < stack.Size-1; i++ {
		newArray[i] = stack.Array[i+1]
	}
	stack.Array = newArray
	stack.Size = stack.Size - 1
	return val
}
