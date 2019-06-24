package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type UpperCaseWriter struct {
	w io.Writer
}

func NewWriter(w io.Writer) io.Writer {
	return UpperCaseWriter{w}
}

// Now UpperCaseWriter implements the io.Writer interface
func (uw UpperCaseWriter) Write(p []byte) (n int, err error) {
	q := strings.ToUpper(string(p))
	return uw.w.Write([]byte(q))
}

func main() {
	w := NewWriter(os.Stdout)
	fmt.Fprintf(w, "hello world")
}
