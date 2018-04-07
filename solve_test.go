package main

import (
	"testing"
)

func TestFitsOnBillboard(t *testing.T) {
	// Test with optimal font size
	listWords := []string{"hacker", "cup"}
	var width, height, fontSize uint = 20, 6, 3

	if !FitsOnBillboard(listWords, width, height, fontSize) {
		t.Fatalf("%v should fit on a %dx%d board with a %d fontSize", listWords, width, height, fontSize)
	}

	// Test with size that is too big
	fontSize = 4
	if FitsOnBillboard(listWords, width, height, fontSize) {
		t.Fatalf("%v can't fit on a %dx%d board with a %d fontSize", listWords, width, height, fontSize)
	}

	// Test with working font size but not optimal
	fontSize = 1
	if !FitsOnBillboard(listWords, width, height, fontSize) {
		t.Fatalf("%v should fit on a %dx%d board with a %d fontSize", listWords, width, height, fontSize)
	}
}
