package reader

import (
	"io"
	"os"
)

func File(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadAll(file *os.File) ([]byte, error) {
	buf := make([]byte, 1024)
	var allBytes []byte
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		allBytes = append(allBytes, buf[:n]...)
	}
	return allBytes, nil
}
