package dll

type DoubleLinkedList[T any] struct {
	length int
	start  *ListElement[T]
	end    *ListElement[T]
}

func NewDLL[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		length: 0,
		start:  nil,
		end:    nil,
	}
}

func (dll *DoubleLinkedList[T]) Add(value T) {
	element := NewDLLElement(value)
	if dll.end == nil {
		dll.start = element
		dll.end = element
	} else {
		end := dll.end

		end.next = element
		element.prev = end

		dll.end = element
	}

	dll.length += 1
}

func (dll *DoubleLinkedList[T]) Insert(selectedElement *ListElement[T], value T) {
	element := NewDLLElement(value)
	next := selectedElement.next

	if next != nil {
		next.prev = element
		element.next = next
	}

	element.prev = selectedElement
	selectedElement.next = element
}

func (dll *DoubleLinkedList[T]) Length() int {
	return dll.length
}

func (dll *DoubleLinkedList[T]) Start() *ListElement[T] {
	return dll.start
}

func (dll *DoubleLinkedList[T]) End() *ListElement[T] {
	return dll.end
}

func (dll *DoubleLinkedList[T]) PrintElements() {
	start := dll.start
	if start != nil {
		println(start.PrintElement())

	}
}
