package heap

func HeapSort(data []interface{}) []interface{} {
	sorts := make([]interface{}, 0)
	heap := &MaxHeap{
		data: data,
	}

	for i := len(data) / 2; i >= 0; i-- {
		heap.siftdown(i)
	}

	for i := 0; i < len(data); i++ {
		sorts = append(sorts, heap.ExtractMin())
	}

	return sorts
}
