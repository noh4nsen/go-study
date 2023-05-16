package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 3)

	for {
		n, err := r.Read(b) // Consome 8 bytes por vez
		fmt.Printf("n = %v; err = %v; b = %v;\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
