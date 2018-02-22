package main

import (
	"os"
	"io"
)

func main() {
	file, err := os.Open(os.Args[0])
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, file)
}
