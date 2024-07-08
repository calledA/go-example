package ds

import (
	"sync"
)

type IArray[T any] interface {
	Len() int
	IsEmpty() bool
	Add(val T) bool
	Concat(val ...[]T) []T
	Remove(fn FliterFunc[T]) bool
	Contain(fn FliterFunc[T]) bool
	Filter(fn FliterFunc[T]) []T
	Map(fn MapFunc[T]) []T
	List() []T
	Clear()
}

func Array[T any](size int) *unsafeArray[T] {
	return newUnsafeArray[T](size)
}

func SafeArray[T any](size int) *safeArray[T] {
	return &safeArray[T]{
		ua: newUnsafeArray[T](size),
	}
}

func newUnsafeArray[T any](size int) *unsafeArray[T] {
	return &unsafeArray[T]{
		array: make([]T, size),
		size:  size,
	}
}

var _ IArray[string] = (*unsafeArray[string])(nil)
var _ IArray[string] = (*safeArray[string])(nil)

// 线程不安全的数组
type unsafeArray[T any] struct {
	array []T
	size  int
}

func (arr *unsafeArray[T]) Add(val T) bool {
	preLen := len(arr.array)
	arr.array = append(arr.array, val)
	arr.size = len(arr.array)
	return arr.size != preLen
}

func (arr *unsafeArray[T]) Append(val ...T) int {
	preLen := len(arr.array)
	arr.array = append(arr.array, val...)
	arr.size = len(arr.array)
	return arr.size - preLen
}

func (arr *unsafeArray[T]) Remove(fn FliterFunc[T]) bool {
	if arr.size == 0 {
		return false
	}

	preLen := len(arr.array)
	newArray := make([]T, 0)
	for _, v := range arr.array {
		if fn(v) {
			continue
		}
		newArray = append(newArray, v)
	}
	arr.array = newArray
	arr.size = len(newArray)
	return preLen != arr.size
}

func (arr *unsafeArray[T]) Contain(fn FliterFunc[T]) bool {
	for _, v := range arr.array {
		if fn(v) {
			return true
		}
	}
	return false
}

func (arr *unsafeArray[T]) Len() int {
	return arr.size
}

func (arr *unsafeArray[T]) IsEmpty() bool {
	return arr.Len() == 0
}

func (arr *unsafeArray[T]) Concat(val ...[]T) []T {
	for _, v := range val {
		arr.array = append(arr.array, v...)
	}
	arr.size = len(arr.array)
	return arr.array
}

func (arr *unsafeArray[T]) Filter(fn FliterFunc[T]) []T {
	newArr := make([]T, 0, arr.size)
	for _, v := range arr.array {
		if fn(v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func (arr *unsafeArray[T]) List() []T {
	return arr.array
}

func (arr *unsafeArray[T]) Map(mf MapFunc[T]) []T {
	newArr := make([]T, 0, arr.size)
	for i, v := range arr.array {
		newArr = append(newArr, mf(i, v))
	}
	return newArr
}

func (arr *unsafeArray[T]) Clear() {
	arr.array = []T{}
	arr.size = 0
}

// 线程不安全的数组
type safeArray[T any] struct {
	ua   *unsafeArray[T]
	lock sync.RWMutex
}

func (arr *safeArray[T]) Add(val T) bool {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Add(val)
}

func (arr *safeArray[T]) Append(val ...T) int {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Append(val...)
}

func (arr *safeArray[T]) Remove(fn FliterFunc[T]) bool {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Remove(fn)
}

func (arr *safeArray[T]) Contain(fn FliterFunc[T]) bool {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Contain(fn)
}

func (arr *safeArray[T]) Len() int {
	return arr.ua.Len()
}

func (arr *safeArray[T]) IsEmpty() bool {
	return arr.ua.IsEmpty()
}

func (arr *safeArray[T]) Concat(val ...[]T) []T {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Concat(val...)
}

func (arr *safeArray[T]) Filter(fn FliterFunc[T]) []T {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Filter(fn)
}

func (arr *safeArray[T]) List() []T {
	return arr.ua.List()
}

func (arr *safeArray[T]) Map(mf MapFunc[T]) []T {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	return arr.ua.Map(mf)
}

func (arr *safeArray[T]) Clear() {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	arr.ua.Clear()
}
