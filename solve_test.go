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

func TestGetRange(t *testing.T) {
	// Test with an original fontsize lower or equal to the optimal fontSize
	listWords := []string{"Facebook", "Hacker", "Cup", "2013"}
	var width, height, fontSize uint = 350, 100, 33

	lowerBound, upperBound := GetRange(listWords, width, height, fontSize)
	if lowerBound != 33 || upperBound != 66 {
		t.Fatalf("Range should be [33, 66], is [%d, %d]", lowerBound, upperBound)
	}

	// Test with an orignal fontsize higher than the optimal fontSize
	fontSize = 34

	lowerBound, upperBound = GetRange(listWords, width, height, fontSize)
	if lowerBound != 17 || upperBound != 34 {
		t.Fatalf("Range should be [17, 34], is [%d, %d]", lowerBound, upperBound)
	}

	// Test with an orignal fontsize far lower than the optimal fontSize
	fontSize = 3

	lowerBound, upperBound = GetRange(listWords, width, height, fontSize)
	if lowerBound != 24 || upperBound != 48 {
		t.Fatalf("Range should be [24, 48], is [%d, %d]", lowerBound, upperBound)
	}

	// Test with an orignal fontsize far higher than the optimal fontSize
	fontSize = 100

	lowerBound, upperBound = GetRange(listWords, width, height, fontSize)
	if lowerBound != 25 || upperBound != 50 {
		t.Fatalf("Range should be [25, 50], is [%d, %d]", lowerBound, upperBound)
	}
}
