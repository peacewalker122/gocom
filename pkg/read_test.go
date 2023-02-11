package pkg

import (
	"gocom/util/heap"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heap.HeapSort([]interface{}{12, 32, 43, 545, 76, 12, 43, 2, 42, 1, 2, 4, 5, 8433247832, 321323, 3123, 122121})
	}
}

func TestExec(t *testing.T) {
}
