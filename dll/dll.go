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

func Of[T any](values ...T) *DoubleLinkedList[T] {
	dll := NewDLL[T]()

	dll.AddAll(values...)

	return dll
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

func (dll *DoubleLinkedList[T]) Insert(selectedElement *ListElement[T], value T) {
	element := NewDLLElement(value) // {element}

	if selectedElement == dll.last {
		dll.last = element
	} else {
		next := selectedElement.next // {selected}-><-{next}

		next.prev = element // {element}<-{next}
		element.next = next // {element}->{next}
		// {element}-><-{next}
	}

	element.prev = selectedElement // {selected}<-{element}
	selectedElement.next = element // {selected}->{element}
	// {selected}-><-{element}-><-{next}

	dll.length += 1
}

func (dll *DoubleLinkedList[T]) InsertAll(selectedElement *ListElement[T], elements *DoubleLinkedList[T]) {
	first := elements.first // {first}-><-...
	last := elements.last   // ...-><-{last}

	if selectedElement == dll.last {
		dll.last = last
	} else {
		next := selectedElement.next // {selected}-><-{next}

		next.prev = last // ...-><-{last}<-{next}
		last.next = next // ...-><-{last}->{next}
		// ...-><-{last}-><-{next}
	}

	first.prev = selectedElement // {selected}<-{first}-><-...
	selectedElement.next = first // {selected}->{first}-><-...
	// {selected}-><-{first}-><-...-><-{last}->{next}

	dll.length += elements.length
}

func (dll *DoubleLinkedList[T]) InsertBefore(selectedElement *ListElement[T], value T) {
	element := NewDLLElement(value)

	if selectedElement == dll.first {
		dll.first = element
	} else {
		prev := selectedElement.prev // {prev}-><-{selected}

		prev.next = element // {prev}->{element}
		element.prev = prev // {prev}<-{element}
		// {prev}-><-{element}
	}

	element.next = selectedElement // {element}->{selected}
	selectedElement.prev = element // {element}<-{selected}
	// {prev}-><-{element}-><-{selected}

	dll.length += 1
}

func (dll *DoubleLinkedList[T]) InsertAllBefore(selectedElement *ListElement[T], elements *DoubleLinkedList[T]) {
	first := elements.first // {first}-><-...
	last := elements.last   // ...-><-{last}

	if selectedElement == dll.first {
		dll.first = first
	} else {
		prev := selectedElement.prev // {prev}-><-{selected}

		prev.next = first // {prev}->{first}-><-...
		first.prev = prev // {prev}<-{first}-><-...
		// {prev}-><-{first}-><-...
	}

	last.next = selectedElement // ...-><-{last}->{selected}
	selectedElement.prev = last // ...-><-{last}<-{selected}
	// {prev}-><-{first}-><-...-><-{last}-><-{selected}

	dll.length += elements.length
}

func (dll *DoubleLinkedList[T]) InsertAt(index int, value T) {
	element := dll.GetIndex(index)
	dll.Insert(element, value)
}

func (dll *DoubleLinkedList[T]) InsertAllAt(index int, values ...T) {
	element := dll.GetIndex(index)
	elements := Of(values...)
	dll.InsertAll(element, elements)
}

func (dll *DoubleLinkedList[T]) InsertBeforeIndex(index int, value T) {
	element := dll.GetIndex(index)
	dll.InsertBefore(element, value)
}

func (dll *DoubleLinkedList[T]) InsertAllBeforeIndex(index int, values ...T) {
	element := dll.GetIndex(index)
	elements := Of(values...)
	dll.InsertAllBefore(element, elements)
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
