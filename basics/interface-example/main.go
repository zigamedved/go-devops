package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	content  string
	position int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.position < len(m.content) {
		n := copy(p, m.content[m.position:m.position+1])
		m.position++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	mySlowReaderInstance := &MySlowReader{content: "Hello World!"}

	out, err := io.ReadAll(mySlowReaderInstance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)
}
