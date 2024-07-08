package ds

import (
	"sync"
)

type Set[T comparable] interface {
	Len() int
	List() []T
	Add(val T) bool
	IsEmpty() bool
	Remove(val T) bool
	Contain(val T) bool
	Append(val ...T) int
	Clear()
}

var _ Set[string] = (*unSafeSet[string])(nil)
var _ Set[string] = (*safeSet[string])(nil)

func NewSet[T comparable]() *unSafeSet[T] {
	return newUnsafeSet[T]()
}

func NewSafeSet[T comparable]() *safeSet[T] {
	return &safeSet[T]{
		uss: newUnsafeSet[T](),
	}
}

func newUnsafeSet[T comparable]() *unSafeSet[T] {
	return &unSafeSet[T]{
		m:    make(map[T]struct{}, 0),
		size: 0,
	}
}

type unSafeSet[T comparable] struct {
	m    map[T]struct{}
	size int
}

func (s *unSafeSet[T]) Add(val T) bool {
	preLen := len(s.m)
	s.m[val] = struct{}{}
	s.size = len(s.m)

	return s.size != preLen
}

func (s *unSafeSet[T]) Append(val ...T) int {
	preLen := len(s.m)
	for _, v := range val {
		s.m[v] = struct{}{}
	}
	s.size = len(s.m)
	return s.size - preLen
}

func (s *unSafeSet[T]) Remove(val T) bool {
	if s.size == 0 {
		return false
	}
	delete(s.m, val)
	s.size = len(s.m)
	return true
}

func (s *unSafeSet[T]) Contain(val T) bool {
	_, ok := s.m[val]
	return ok
}

func (s *unSafeSet[T]) Len() int {
	return s.size
}

func (s *unSafeSet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *unSafeSet[T]) Clear() {
	s.m = map[T]struct{}{}
	s.size = 0
}

func (s *unSafeSet[T]) List() []T {
	list := make([]T, s.size)
	for val := range s.m {
		list = append(list, val)
	}
	return list
}

type safeSet[T comparable] struct {
	uss  *unSafeSet[T]
	lock sync.RWMutex
}

func (s *safeSet[T]) Add(val T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.uss.Add(val)
}

func (s *safeSet[T]) Append(val ...T) int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.uss.Append(val...)
}

func (s *safeSet[T]) Remove(val T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.uss.Remove(val)
}

func (s *safeSet[T]) Contain(val T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.uss.Contain(val)
}

func (s *safeSet[T]) Len() int {
	return s.uss.Len()
}

func (s *safeSet[T]) IsEmpty() bool {
	return s.uss.IsEmpty()
}

func (s *safeSet[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.uss.Clear()
}

func (s *safeSet[T]) List() []T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.uss.List()
}
