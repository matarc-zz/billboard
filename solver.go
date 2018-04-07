package main

import (
	"fmt"
	"io"
	"os"
)

func Solver(reader io.Reader, writer io.Writer) {
	var numberOfLines int
	var err error
	r := NewParser(reader)
	if numberOfLines, err = r.ParseNumberOfLines(); err != nil {
		return
	}
	var width, height uint64
	var listWords []string
	var fontSize uint
	for i := 1; i <= numberOfLines; i++ {
		if listWords, width, height, err = r.ParseLine(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fontSize = Solve(listWords, uint(width), uint(height))
		if _, err = writer.Write([]byte(fmt.Sprintf("Case #%d: %d\n", i, fontSize))); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
