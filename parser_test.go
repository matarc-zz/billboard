package main

import (
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "20 6 hacker cup"
	p := NewParser(strings.NewReader(line))
	listWords, width, height, err := p.ParseLine()
	if err != nil {
		t.Fatal(err)
	}
	if width != 20 {
		t.Fatalf("Width should be 20, is %d", width)
	}
	if height != 6 {
		t.Fatalf("Height should be 6, is %d", height)
	}
	if len(listWords) != 2 {
		t.Fatalf("len(listWords) should be 2, is %d", len(listWords))
	}
	if listWords[0] != "hacker" {
		t.Fatalf("First word should be hacker, is %s", listWords[0])
	}
	if listWords[1] != "cup" {
		t.Fatalf("First word should be cup, is %s", listWords[0])
	}
}

func TestParseNumberOfLines(t *testing.T) {
	line := "22"
	p := NewParser(strings.NewReader(line))
	numberOfLines, err := p.ParseNumberOfLines()
	if err != nil {
		t.Fatal(err)
	}
	if numberOfLines != 22 {
		t.Fatalf("Number of lines should be 22, is %d", numberOfLines)
	}
}
