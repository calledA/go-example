package ds

import (
	"testing"
)

func TestUnsafeArray(t *testing.T) {
	ua := Array[string](0)

	t.Log(ua.IsEmpty())
	t.Log(ua.Add("test1"))
	t.Log(ua.Append("test2", "test3"))

	t.Log(ua.Contain(func(item string) bool {
		return item == "test4"
	}))

	t.Log(ua.Remove(func(item string) bool {
		return item == "test2"
	}))

	t.Log(ua.Map(func(index int, item string) string {
		return item + "test"
	}))

	t.Log(ua.List())
	t.Log(ua.IsEmpty())

	ua.Clear()
	t.Log(ua.IsEmpty())
}
