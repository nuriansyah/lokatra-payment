package shared_test

import (
	"testing"

	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/stretchr/testify/assert"
)

func TestBatching(t *testing.T) {
	t.Run("success batch 1 over 5", func(t *testing.T) {
		b := shared.NewBatch(1, []int{1, 2, 3, 4, 5})
		actual := b.MustBatchingReq()
		expected := []interface{}{
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{4},
			[]int{5},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("success batch 3 over 5", func(t *testing.T) {
		b := shared.NewBatch(3, []int{1, 2, 3, 4, 5})
		actual := b.MustBatchingReq()
		expected := []interface{}{
			[]int{1, 2, 3},
			[]int{4, 5},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("success batch 5 over 5", func(t *testing.T) {
		b := shared.NewBatch(5, []int{1, 2, 3, 4, 5})
		actual := b.MustBatchingReq()
		expected := []interface{}{
			[]int{1, 2, 3, 4, 5},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("success batch 6 over 5", func(t *testing.T) {
		b := shared.NewBatch(6, []int{1, 2, 3, 4, 5})
		actual := b.MustBatchingReq()
		expected := []interface{}{
			[]int{1, 2, 3, 4, 5},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestBatchingSize(t *testing.T) {
	t.Run("success batch size 5 over 19", func(t *testing.T) {
		b := shared.NewBatch(19, nil)
		actual := b.BatchingSize(5)
		expected := []int{5, 5, 5, 4}

		assert.Equal(t, expected, actual)
	})

	t.Run("success batch size 5 over 20", func(t *testing.T) {
		b := shared.NewBatch(20, nil)
		actual := b.BatchingSize(5)
		expected := []int{5, 5, 5, 5}

		assert.Equal(t, expected, actual)
	})

	t.Run("success batch size 5 over 2", func(t *testing.T) {
		b := shared.NewBatch(2, nil)
		actual := b.BatchingSize(5)
		expected := []int{2}

		assert.Equal(t, expected, actual)
	})
}
