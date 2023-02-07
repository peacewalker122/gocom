package heap

import (
	"gocom/entity"
)

type MinHeap struct {
	data []interface{}
}

func NewMinHeap() *MinHeap {
	data := make([]interface{}, 0)
	return &MinHeap{data: data}
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Less(i, j int) bool {
	iData := h.data[i].(*entity.Tree)
	jData := h.data[j].(*entity.Tree)
	return iData.Freq > jData.Freq
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.data = append(h.data, x)
	h.siftup(len(h.data) - 1)
}

func (h *MinHeap) Pop() interface{} {
	if len(h.data) == 0 {
		return nil
	}
	h.Swap(0, len(h.data)-1)
	x := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftdown(0)
	return x
}

func (h *MinHeap) UpdateByIndex(index int, newData interface{}) {
	old := h.data[index].(*entity.Tree)
	h.data[index] = newData.(*entity.Tree)
	if old.Freq > newData.(*entity.Tree).Freq {
		h.siftup(index)
	} else {
		h.siftdown(index)
	}
}

func (h *MinHeap) UpdateByValue(oldData interface{}, newData interface{}) {
	for index, data := range h.data {
		if data == oldData {
			h.UpdateByIndex(index, newData)
			return
		}
	}
}

func (h *MinHeap) Top() interface{} {
	return h.data[0]
}

func (h *MinHeap) siftdown(index int) {
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

func (h *MinHeap) siftup(index int) {
	parent := (index - 1) / 2
	if index > 0 && h.Less(parent, index) {
		h.Swap(parent, index)
		h.siftup(parent)
	} else {
		return
	}
}
