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
		numbers := dll.Of(23, 45, 37)

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

	t.Run("insert", func(t *testing.T) {
		numbers := dll.Of(5, 7)

		numbers.InsertAt(0, 6)
		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
	})

	t.Run("insert at end", func(t *testing.T) {
		numbers := dll.Of(5, 6)

		numbers.InsertAt(1, 7)
		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 7, numbers.GetLast().GetValue())
	})

	t.Run("insert all", func(t *testing.T) {
		numbers := dll.Of(5, 9)

		numbers.InsertAllAt(0, 6, 7, 8)
		assert.Equal(t, 5, numbers.Length())

		assert.Equal(t, 5, numbers.GetFirst().GetValue())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
		assert.Equal(t, 7, numbers.GetIndex(2).GetValue())
		assert.Equal(t, 8, numbers.GetIndex(3).GetValue())
		assert.Equal(t, 9, numbers.GetLast().GetValue())
	})

	t.Run("insert all at end", func(t *testing.T) {
		numbers := dll.Of(5, 6)

		numbers.InsertAllAt(1, 7, 8, 9)
		assert.Equal(t, 5, numbers.Length())

		assert.Equal(t, 5, numbers.GetFirst().GetValue())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
		assert.Equal(t, 7, numbers.GetIndex(2).GetValue())
		assert.Equal(t, 8, numbers.GetIndex(3).GetValue())
		assert.Equal(t, 9, numbers.GetLast().GetValue())
	})

	t.Run("insert before", func(t *testing.T) {
		numbers := dll.Of(5, 7)

		numbers.InsertBeforeIndex(1, 6)

		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
	})

	t.Run("insert before first", func(t *testing.T) {
		numbers := dll.Of(5, 6)

		numbers.InsertBeforeIndex(0, 4)

		assert.Equal(t, 3, numbers.Length())
		assert.Equal(t, 4, numbers.GetFirst().GetValue())
	})

	t.Run("insert all before", func(t *testing.T) {
		numbers := dll.Of(5, 9)

		numbers.InsertAllBeforeIndex(1, 6, 7, 8)
		assert.Equal(t, 5, numbers.Length())

		assert.Equal(t, 5, numbers.GetFirst().GetValue())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
		assert.Equal(t, 7, numbers.GetIndex(2).GetValue())
		assert.Equal(t, 8, numbers.GetIndex(3).GetValue())
		assert.Equal(t, 9, numbers.GetLast().GetValue())
	})

	t.Run("insert all before first", func(t *testing.T) {
		numbers := dll.Of(8, 9)

		numbers.InsertAllBeforeIndex(0, 5, 6, 7)
		assert.Equal(t, 5, numbers.Length())

		assert.Equal(t, 5, numbers.GetFirst().GetValue())
		assert.Equal(t, 6, numbers.GetIndex(1).GetValue())
		assert.Equal(t, 7, numbers.GetIndex(2).GetValue())
		assert.Equal(t, 8, numbers.GetIndex(3).GetValue())
		assert.Equal(t, 9, numbers.GetLast().GetValue())
	})
}
