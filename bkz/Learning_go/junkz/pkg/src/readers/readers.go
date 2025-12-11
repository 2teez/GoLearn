package readers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	n, _, err := reader.ReadLine()
	if err != nil {
		return string([]byte{}), err
	}
	return string(n), nil
}

func GetFloat(msg string) (float64, error) {
	fmt.Printf("%s", msg)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
