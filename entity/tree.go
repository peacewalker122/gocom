package entity

type Tree struct {
	Bytes byte
	Freq  int
	Left  *Tree
	Right *Tree
}

func (t Tree) Less(other Tree) bool {
	return t.Freq < other.Freq
}
