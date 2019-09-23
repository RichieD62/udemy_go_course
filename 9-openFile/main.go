package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {
	txtFile := os.Args[1]

	opener, err := os.Open(txtFile)
	if err != nil {
		errorMsg := fmt.Sprintf("failed to open txt file - error: %v", err)
		log.Printf(errorMsg)
	}

	lw := logWriter{}

	io.Copy(lw, opener)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
