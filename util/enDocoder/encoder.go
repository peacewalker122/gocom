package enDocoder

import (
	"encoding/binary"
	"encoding/xml"
	"gocom/entity"
	"os"
)

const (
	headerSize  = 24
	versionByte = 1
)

type header struct {
	Version   uint8    `gob:"Version"`
	DataSize  int64    `gob:"DataSize"`
	Encoding  uint8    `gob:"Encoding"`
	Reserved1 [7]byte  `gob:"Reserved1"`
	Reserved2 [16]byte `gob:"Reserved2"`
}

func writeHeader(file *os.File, dataSize int64) error {
	var h header

	h.Version = versionByte
	h.DataSize = dataSize
	h.Encoding = 1

	return binary.Write(file, binary.LittleEndian, &h)
}

func WriteEncodedDataToFile(encoded []byte, tree *entity.Tree, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create an XML encoder
	encoder := xml.NewEncoder(file)

	//err = writeHeader(file, int64(len(encoded)))
	//if err != nil {
	//	return err
	//}

	// Write the encoded data to the file
	_, err = file.Write(encoded)
	if err != nil {
		return err
	}

	// Encode the tree to the file
	return encoder.Encode(tree)
}
