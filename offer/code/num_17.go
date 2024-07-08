package code

import (
	"gmr/gmr-x/ds"
)

func Num17MethodOne(a, b *ds.BTree) bool {
	if a == nil || b == nil {
		return false
	}
	return isSub(a, b)
}

func isSub(a, b *ds.BTree) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Data != b.Data {
		return false
	}
	return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}
