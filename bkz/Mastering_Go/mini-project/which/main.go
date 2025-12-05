package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Provide file name..")
		return
	}

	for _, file := range args[1:] {
		path := os.Getenv("PATH")
		for _, dir := range filepath.SplitList(path) {
			fullPath := filepath.Join(dir, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}
