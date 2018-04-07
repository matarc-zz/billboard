package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestSolver(t *testing.T) {
	input := `5
		  20 6 hacker cup
		  100 20 hacker cup 2013
		  10 20 MUST BE ABLE TO HACK
		  55 25 Can you hack
		  100 20 Hack your way to the cup`
	output := `Case #1: 3
Case #2: 10
Case #3: 2
Case #4: 8
Case #5: 7`
	reader := strings.NewReader(input)
	writer := new(bytes.Buffer)
	Solver(reader, writer)
	scannerResult := bufio.NewScanner(writer)
	scannerMock := bufio.NewScanner(strings.NewReader(output))
	for scannerResult.Scan() && scannerMock.Scan() {
		if err := scannerResult.Err(); err != nil {
			t.Fatal(err)
		}
		if err := scannerMock.Err(); err != nil {
			t.Fatal(err)
		}
		if scannerResult.Text() != scannerMock.Text() {
			t.Fatalf("Should be `%s`, is `%s`", scannerMock.Text(), scannerResult.Text())
		}
	}
}
