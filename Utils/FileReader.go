package Utils

import (
	"io"
	"os"
)

func ReadFileToByte(path string) []byte {

	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	return data

}
