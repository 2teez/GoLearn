package readers

import (
	"bufio"
	"os"
)

func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	n, _, err := reader.ReadLine()
	if err != nil {
		return string([]byte{}), err
	}
	return string(n), nil
}
