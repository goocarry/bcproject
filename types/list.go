package types

import (
	"fmt"
	"reflect"
)

// List is a generic collection.
type List[T any] struct {
	Data []T
}

// NewList returns new generic collection.
func NewList[T any]() *List[T] {
	return &List[T]{
		Data: []T{},
	}
}

func (l *List[T]) Get(index int) T {
	if index > len(l.Data) {
		err := fmt.Sprintf("the given index (%d) is higher than the length (%d) of the list", index, len(l.Data))
		panic(err)
	}
	return l.Data[index]
}

func (l *List[T]) Insert(element T) {
	l.Data = append(l.Data, element)
}

func (l *List[T]) Clear() {
	l.Data = []T{}
}

// GetIndex will return the index of v. If v does not exist in the list
// -1 will be returned.
func (l *List[T]) GetIndex(element T) int {
	for i := range l.Data {
		if reflect.DeepEqual(l.Data[i], element) {
			return i
		}
	}
	return -1
}

func (l *List[T]) Remove(element T) {
	index := l.GetIndex(element)
	if index == -1 {
		return
	}
	l.Pop(index)
}

func (l *List[T]) Pop(index int) {
	l.Data = append(l.Data[:index], l.Data[index+1:]...)
}

func (l *List[T]) Contains(v T) bool {
	for i := 0; i < len(l.Data); i++ {
		if reflect.DeepEqual(l.Data[i], v) {
			return true
		}
	}
	return false
}

func (l List[T]) Last() T {
	return l.Data[l.Len()-1]
}

func (l *List[T]) Len() int {
	return len(l.Data)
}
