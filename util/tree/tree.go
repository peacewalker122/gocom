package tree

import (
	"gocom/entity"
	"gocom/util/counter"
	"gocom/util/queue"
)

func BuildTree(bytes []byte) *entity.Tree {
	counter := counter.NewMaps(bytes)

	pq := queue.NewQueue()
	for bytes, freq := range counter {
		pq.EnQueue(&entity.Tree{
			Bytes: bytes,
			Freq:  freq,
		})
	}

	for pq.Len() > 1 {
		var left, right *entity.Tree
		var freqSum int
		if n, ok := pq.DeQueue().(*entity.Tree); ok {
			left = n
		}
		if n, ok := pq.DeQueue().(*entity.Tree); ok {
			right = n
		}

		if left != nil {
			freqSum += left.Freq
		}
		if right != nil {
			freqSum += right.Freq
		}

		parent := &entity.Tree{
			Freq:  freqSum,
			Left:  left,
			Right: right,
		}
		pq.EnQueue(parent)
	}

	return pq.DeQueue().(*entity.Tree)
}

func BuildMap(root *entity.Tree) map[byte]string {
	var dfs func(root *entity.Tree, code string, encodingMap map[byte]string)
	dfs = func(root *entity.Tree, code string, encodingMap map[byte]string) {
		if root.Bytes != 0 {
			encodingMap[root.Bytes] = code
			return
		}
		dfs(root.Left, code+"0", encodingMap)
		dfs(root.Right, code+"1", encodingMap)
	}
	encodingMap := make(map[byte]string)
	dfs(root, "", encodingMap)
	return encodingMap
}

func Encode(bytes []byte) ([]byte, *entity.Tree) {
	root := BuildTree(bytes)
	encodingMap := BuildMap(root)

	var encoded []byte
	for _, b := range bytes {
		encoded = append(encoded, encodingMap[b]...)
	}

	return encoded, root
}

func Decode(encoded []byte, tree *entity.Tree) string {
	var decoded []byte
	var root *entity.Tree
	root = tree
	for _, b := range encoded {
		if b == '0' {
			root = root.Left
		} else {
			root = root.Right
		}

		if root.Bytes != 0 {
			decoded = append(decoded, root.Bytes)
			root = tree
		}
	}
	return string(decoded)
}
