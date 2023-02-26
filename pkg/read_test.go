package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gocom/util/enDocoder"
	util "gocom/util/fileSystem"
	"gocom/util/heap"
	"gocom/util/tree"
	"io"
	"log"
	"os"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heap.HeapSort([]interface{}{12, 32, 43, 545, 76, 12, 43, 2, 42, 1, 2, 4, 5, 8433247832, 321323, 3123, 122121})
	}
}

func TestEncodeAndDecode(t *testing.T) {
	util.WalkDirectory("../")
	file, err := os.ReadFile("assets/file.txt")
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	encodeByte, encodeTree := tree.Encode(file)
	log.Println(encodeByte)
	err = enDocoder.WriteEncodedDataToFile(encodeByte, encodeTree, "assets/file.txt.huff")
	require.NoError(t, err)

	bytes, treeData, err := enDocoder.ReadEncodedDataFromFile("assets/file.txt.huff")
	require.NoError(t, err)

	assert.Equal(t, encodeTree, treeData)

	log.Println(tree.Decode(bytes, treeData))
}

func TestExec(t *testing.T) {
	log.Println(os.Getwd())

	log.Println(util.WalkDirectory("/usr"))

	log.Println(os.Getwd())
}
