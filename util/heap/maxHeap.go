package heap

type MaxHeap struct {
	data []interface{}
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.data[i].(int) > h.data[j].(int)
}

func (h *MaxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Push(x interface{}) {
	h.data = append(h.data, x)
	h.siftup(len(h.data) - 1)
}

func (h *MaxHeap) Pop() interface{} {
	if len(h.data) == 0 {
		return nil
	}
	h.Swap(0, len(h.data)-1)
	x := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftdown(0)
	return x
}

func (h *MaxHeap) Top() interface{} {
	return h.data[0]
}

func (h *MaxHeap) UpdateByIndex(index int, newData interface{}) {
	old := h.data[index]
	h.data[index] = newData
	if old.(int) > newData.(int) {
		h.siftup(index)
	} else {
		h.siftdown(index)
	}
}

func (h *MaxHeap) UpdateByValue(oldData interface{}, newData interface{}) {
	for index, data := range h.data {
		if data == oldData {
			h.UpdateByIndex(index, newData)
			return
		}
	}
}

func (s *MaxHeap) ExtractMin() interface{} {
	return s.Pop()
}

func (h *MaxHeap) siftdown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index
	if left < len(h.data) && h.Less(largest, left) {
		largest = left
	}
	if right < len(h.data) && h.Less(largest, right) {
		largest = right
	}
	if largest != index {
		h.Swap(index, largest)
		h.siftdown(largest)
	} else {
		return
	}
}

func (h *MaxHeap) siftup(index int) {
	parent := (index - 1) / 2
	if index > 0 && h.Less(parent, index) {
		h.Swap(parent, index)
		h.siftup(parent)
	} else {
		return
	}
}
