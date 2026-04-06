package addition

import (
	"fmt"
	"strconv"
)

func Sum(start int, o []string) int {
	result := start
	for _, value := range o {
		value, err := strconv.Atoi(value)
		if err != nil {
			defer func() {
				fmt.Println(fmt.Errorf("%v", err))
			}()
		}
		result += value
	}
	return result
}
