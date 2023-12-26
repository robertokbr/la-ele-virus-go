package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	msg := "La ele!"

	for i := 0; i < 1000; i++ {
		file, _ := os.Create(fmt.Sprintf("la-ele-%d.txt", i))
		data := bytes.NewBuffer([]byte(msg))
		defer file.Close()
		io.Copy(file, data)
	}
}
