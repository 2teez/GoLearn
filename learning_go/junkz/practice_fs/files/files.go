package files

import "os"

func Create(filename string) (*os.File, func(), error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	return file, func() { file.Close() }, nil
}

func Write(file *os.File, data []byte) (int, error) {
	return file.Write(data)
}
