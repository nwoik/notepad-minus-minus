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

func (dll *DoubleLinkedList[T]) AddAll(values ...T) {
	for _, value := range values {
		dll.Add(value)
	}
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

	dll.length += 1
}

func (dll *DoubleLinkedList[T]) InsertAtIndex(index int, value T) {
	element := dll.GetIndex(index)
	dll.Insert(element, value)
}

func (dll *DoubleLinkedList[T]) Replace(index int, value T) {
	element := dll.GetIndex(index)
	element.SetValue(value)
}

func (dll *DoubleLinkedList[T]) GetIndex(index int) *ListElement[T] {
	if index >= dll.Length() {
		return dll.end
	}

	count := 0
	element := dll.Start()

	for count < dll.Length() {
		if count == index {
			return element
		}
		element = element.next
		count++
	}

	return dll.end
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
