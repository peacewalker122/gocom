package pkg

import (
	"fmt"
	"gocom/util/tree"
	"io"
	"log"
	"os"
	"time"
)

func Exec() {
	file, err := os.ReadFile("assets/file.txt")
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	encodeByte, encodeTree := tree.Encode(file)

	decodedByte := tree.Decode(encodeByte, encodeTree)

	log.Println(string(decodedByte))

	select {
	case <-time.After(200 * time.Millisecond):
		return
	}
}
