package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func CreateKeyByConcatenation(a, b, c string) string {
	return a + ":" + b + ":" + c
}

func CreateKeyByBytesBuffer(a, b, c string) string {
	l := len(a) + len(b) + len(c) + 2
	buf := make([]byte, 0, l)
	w := bytes.NewBuffer(buf)
	w.WriteString(a)
	w.WriteRune(':')
	w.WriteString(b)
	w.WriteRune(':')
	w.WriteString(c)
	return w.String()
}

func CreateKeyByBytesBufferWithoutLen(items ...string) string {
	var buf bytes.Buffer
	for i := range items {
		buf.WriteString(items[i])
		if i < len(items)-1 {
			buf.WriteRune(':')
		}
	}
	return buf.String()
}
