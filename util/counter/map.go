package counter

import (
	"fmt"
	"strings"
)

func (m Maps) Add(b byte) {
	m[b]++
}

func (m Maps) Get(b byte) int {
	return m[b]
}

func (m Maps) String() string {
	var strs []string

	for k, v := range m {
		strs = append(strs, fmt.Sprintf("%d: %d", k, v))
	}

	return strings.Join(strs, "\n")
}

type Maps map[byte]int

func NewMaps(bytes []byte) Maps {
	data := make(Maps)
	for _, b := range bytes {
		data.Add(b)
	}

	return data
}
