package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Println("Specify a file to use.")
	}
	filename := args[1]
	f, closer, err := getFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	words := make([]byte, 2084)
	for {
		n, err := f.Read(words)
		os.Stdout.Write(words[:n])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func getFile(file string) (*os.File, func(), error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		f.Close()
	}, nil
}
