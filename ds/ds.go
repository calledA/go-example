package ds

// 过滤回调
type FliterFunc[T any] func(item T) bool

// 遍历回调
type MapFunc[T any] func(index int, item T) T
