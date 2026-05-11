package dll

import "fmt"

type ListElement[T any] struct {
	value T
	prev  *ListElement[T]
	next  *ListElement[T]
}

func NewDLLElement[T any](value T) *ListElement[T] {
	return &ListElement[T]{
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (element *ListElement[T]) GetValue() T {
	return element.value
}

func (element *ListElement[T]) SetValue(value T) {
	element.value = value
}

func (element *ListElement[T]) PrintElement() string {

	next := element.next
	var strNext string

	if next != nil {
		strNext = next.PrintElement()
	}

	return fmt.Sprintf("<-%v->%s", element.value, strNext)
}
