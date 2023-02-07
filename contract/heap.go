package contract

type Heap interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
	Top() interface{}
	UpdateByIndex(index int, newData interface{})
	UpdateByValue(oldData interface{}, newData interface{})
}
