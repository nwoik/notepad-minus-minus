package tests

import (
	"notepad-minus-minus/dll"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("add 1 number", func(t *testing.T) {
		numbers := dll.NewDLL[int]()
		numbers.Add(23)

		assert.Equal(t, 1, numbers.Length())
		assert.Equal(t, numbers.GetFirst(), numbers.GetLast())
	})

	t.Run("add 3 numbers", func(t *testing.T) {
		numbers := dll.NewDLL[int]()
		numbers.Add(23)
		numbers.Add(45)
		numbers.Add(37)

		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 23, numbers.GetFirst().GetValue())
		assert.Equal(t, 37, numbers.GetLast().GetValue())

	})

	t.Run("add array of numbers", func(t *testing.T) {
		numbers := dll.NewDLL[int]()
		array := []int{1, 2, 3}

		numbers.AddAll(array...)
		assert.Equal(t, 3, numbers.Length())
	})

	t.Run("get indices", func(t *testing.T) {
		numbers := dll.NewDLL[int]()
		numbers.Add(23)
		numbers.Add(45)
		numbers.Add(37)

		assert.Equal(t, 23, numbers.GetIndex(0).GetValue())
		assert.Equal(t, 45, numbers.GetIndex(1).GetValue())
		assert.Equal(t, 37, numbers.GetIndex(2).GetValue())
	})

	t.Run("replace number", func(t *testing.T) {
		numbers := dll.NewDLL[int]()

		numbers.Add(5)
		numbers.Replace(0, 2)
		assert.Equal(t, 2, numbers.GetFirst().GetValue())
	})

	t.Run("insert after", func(t *testing.T) {
		numbers := dll.NewDLL[int]()

		element := numbers.Add(5)
		numbers.Add(7)

		numbers.InsertAfter(element, 6)
		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
	})

	t.Run("insert before", func(t *testing.T) {
		numbers := dll.NewDLL[int]()

		element := numbers.Add(5)
		numbers.Add(6)

		numbers.InsertBefore(element, 4)

		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 4, numbers.GetFirst().GetValue())
	})
}
