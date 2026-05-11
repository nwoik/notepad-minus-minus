package dll

type DoubleLinkedList[T any] struct {
	length int
	first  *ListElement[T]
	last   *ListElement[T]
}

func NewDLL[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		length: 0,
		first:  nil,
		last:   nil,
	}
}

func (dll *DoubleLinkedList[T]) Add(value T) *ListElement[T] {
	element := NewDLLElement(value)
	if dll.last == nil {
		dll.first = element
		dll.last = element
	} else {
		end := dll.last

		end.next = element
		element.prev = end

		dll.last = element
	}

	dll.length += 1
	return element
}

func (dll *DoubleLinkedList[T]) AddAll(values ...T) {
	for _, value := range values {
		dll.Add(value)
	}
}

func (dll *DoubleLinkedList[T]) InsertAfter(selectedElement *ListElement[T], value T) {
	element := NewDLLElement(value)
	next := selectedElement.next

	if next != nil {
		next.prev = element
		element.next = next
	}

	element.prev = selectedElement
	selectedElement.next = element

	if selectedElement == dll.last {
		dll.last = element
	}
	dll.length += 1
}

func (dll *DoubleLinkedList[T]) InsertBefore(selectedElement *ListElement[T], value T) {
	element := NewDLLElement(value)
	prev := selectedElement.prev

	if prev != nil {
		prev.next = element
		element.prev = prev
	}

	element.next = selectedElement
	selectedElement.prev = element

	if selectedElement == dll.first {
		dll.first = element
	}

	dll.length += 1
}

func (dll *DoubleLinkedList[T]) InsertAfterIndex(index int, value T) {
	element := dll.GetIndex(index)
	dll.InsertAfter(element, value)
}

func (dll *DoubleLinkedList[T]) InsertBeforeIndex(index int, value T) {
	element := dll.GetIndex(index)
	dll.InsertBefore(element, value)
}

func (dll *DoubleLinkedList[T]) Replace(index int, value T) {
	element := dll.GetIndex(index)
	element.SetValue(value)
}

func (dll *DoubleLinkedList[T]) GetIndex(index int) *ListElement[T] {
	if index >= dll.Length() {
		return dll.last
	}

	count := 0
	element := dll.GetFirst()

	for count < dll.Length() {
		if count == index {
			return element
		}
		element = element.next
		count++
	}

	return dll.last
}

func (dll *DoubleLinkedList[T]) Length() int {
	return dll.length
}

func (dll *DoubleLinkedList[T]) GetFirst() *ListElement[T] {
	return dll.first
}

func (dll *DoubleLinkedList[T]) GetLast() *ListElement[T] {
	return dll.last
}

func (dll *DoubleLinkedList[T]) PrintElements() {
	start := dll.first
	if start != nil {
		println(start.PrintElement())
	}
}
