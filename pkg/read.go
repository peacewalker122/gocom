package pkg

import (
	"fmt"
	"gocom/util/enDocoder"
	"gocom/util/tree"
	"io"
	"log"
	"os"
	"time"
)

func Create(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	encodeByte, encodeTree := tree.Encode(file)

	err = enDocoder.WriteEncodedDataToFile(encodeByte, encodeTree, "assets/file.txt.huff")
	if err != nil {
		panic(err)
	}
}

func ExecFile() {
	bytes, treeData, err := enDocoder.ReadEncodedDataFromFile("assets/file.txt.huff")
	if err != nil {
		panic(err)
	}

	log.Println(tree.Decode(bytes, treeData))

	select {
	case <-time.After(200 * time.Millisecond):
		return
	}
}
