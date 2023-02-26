package enDocoder

import (
	"bytes"
	"encoding/xml"
	"gocom/entity"
	"gocom/util/helper"
	"io"
	"log"
	"os"
)

func ReadEncodedDataFromFile(filePath string) ([]byte, *entity.Tree, error) {
	tree := &entity.Tree{}
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

	err = decoder.Decode(tree)
	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	_, endIndex := helper.FindValidRangeASCIIBinary(encodedData)
	log.Println(endIndex)
	return encodedData[0:endIndex], tree, nil
}
