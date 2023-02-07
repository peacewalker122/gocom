package pkg

import (
	"gocom/util/heap"
	"log"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exec()
	}
}

func TestExec(t *testing.T) {
	sort := heap.HeapSort([]interface{}{12, 32, 43, 545, 76, 12, 43, 2, 42, 1, 2, 4, 5})

	log.Println(sort)
}
