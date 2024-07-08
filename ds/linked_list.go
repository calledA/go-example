package ds

import "sync"

type ILinkedList[T any] interface {
	Len() int
	Add(val T) bool
	IsEmpty() bool
	Remove(fn FliterFunc[T]) bool
	Contain(fn FliterFunc[T]) bool
	First() ILinkedList[T]
	Last() ILinkedList[T]
	Clear()
}

var _ ILinkedList[string] = (*unsafeLinkedList[string])(nil)
var _ ILinkedList[string] = (*safeLinkedList[string])(nil)

type unsafeLinkedList[T any] struct {
	Val  T
	Next *unsafeLinkedList[T]
}

func (ln *unsafeLinkedList[T]) First() ILinkedList[T] {
	if ln.Next == nil {
		return nil
	}
	return ln.Next
}

func (ln *unsafeLinkedList[T]) Last() ILinkedList[T] {
	if ln.Next == nil {
		return nil
	}

	n := ln
	for n.Next != nil {
		n = n.Next
		if n.Next == nil {
			return n
		}
	}
	return nil
}

func (ln *unsafeLinkedList[T]) Add(val T) bool {

	n := new(unsafeLinkedList[T])
	n.Val = val
	ln.Next = n
	return true
}

func (ln *unsafeLinkedList[T]) Remove(fn FliterFunc[T]) bool {
	if ln.Next == nil {
		return false
	}

	for ln.Next != nil {
		if fn(ln.Next.Val) {
			ln.Next = ln.Next.Next
			return true
		}
		ln = ln.Next
	}
	return false
}

func (ln *unsafeLinkedList[T]) Contain(fn FliterFunc[T]) bool {
	n := ln
	for n.Next != nil {
		if fn(n.Val) {
			return true
		}
		n = n.Next
	}
	return false
}

func (ln *unsafeLinkedList[T]) Len() int {
	len := 0
	n := ln
	for n.Next != nil {
		len++
		n = n.Next
	}
	return len
}

func (ln *unsafeLinkedList[T]) IsEmpty() bool {
	return ln.Next == nil
}

func (ln *unsafeLinkedList[T]) Clear() {
	if ln.Next == nil {
		return
	}
	for ln.Next != nil {
		next := ln.Next
		ln.Next = nil
		ln = next
	}
}

type safeLinkedList[T any] struct {
	ull  unsafeLinkedList[T]
	lock sync.RWMutex
}

func (ln *safeLinkedList[T]) First() ILinkedList[T] {
	return ln.ull.First()
}

func (ln *safeLinkedList[T]) Last() ILinkedList[T] {
	ln.lock.Lock()
	defer ln.lock.Unlock()

	return ln.ull.Last()
}

func (ln *safeLinkedList[T]) Add(val T) bool {
	ln.lock.Lock()
	defer ln.lock.Unlock()

	return ln.ull.Add(val)
}

func (ln *safeLinkedList[T]) Remove(fn FliterFunc[T]) bool {
	ln.lock.Lock()
	defer ln.lock.Unlock()

	return ln.ull.Remove(fn)
}

func (ln *safeLinkedList[T]) Contain(fn FliterFunc[T]) bool {
	ln.lock.RLock()
	defer ln.lock.RUnlock()

	return ln.ull.Contain(fn)
}

func (ln *safeLinkedList[T]) Len() int {
	return ln.ull.Len()
}

func (ln *safeLinkedList[T]) IsEmpty() bool {
	return ln.ull.IsEmpty()
}

func (ln *safeLinkedList[T]) Clear() {
	ln.lock.Lock()
	defer ln.lock.Unlock()

	ln.ull.Clear()
}
