package enDocoder

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"gocom/entity"
	"io"
	"os"
)

func ReadEncodedDataFromFile(filePath string) ([]byte, *entity.Tree, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Read the encoded data from the file
	encodedData, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	// Decode the tree from the encoded data
	buf := bytes.NewBuffer(encodedData)
	decoder := xml.NewDecoder(buf)
	tree := &entity.Tree{}
	err = decoder.Decode(tree)
	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	return encodedData[:len(encodedData)-binary.Size(tree)], tree, nil
}
