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

	// Test when the words can't fit on the billboard no matter which fontsize you use
	width, height, fontSize = 3, 100, 33

	lowerBound, upperBound = GetRange(listWords, width, height, fontSize)
	if lowerBound != 0 || upperBound != 1 {
		t.Fatalf("Range should be [0, 1], is [%d, %d]", lowerBound, upperBound)
	}
}

func TestMin(t *testing.T) {
	if Min(2, 5) != 2 {
		t.Fatalf("Min(2, 5) should be 2")
	}
	if Min(5, 2) != 2 {
		t.Fatalf("Min(5, 2) should be 2")
	}
}

func TestMax(t *testing.T) {
	if Max(2, 5) != 5 {
		t.Fatalf("Max(2, 5) should be 5")
	}
	if Max(5, 2) != 5 {
		t.Fatalf("Max(5, 2) should be 5")
	}
}

func TestSolve(t *testing.T) {
	// Test when width > height
	listWords := []string{"Facebook", "Hacker", "Cup", "2013"}
	var width, height uint = 350, 100

	fontSize := Solve(listWords, width, height)
	if fontSize != 33 {
		t.Fatalf("%v should fit on a %dx%d board with a 33 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

	// Test when width < height
	listWords = []string{"MUST", "BE", "ABLE", "TO", "HACK"}
	width, height = 10, 20

	fontSize = Solve(listWords, width, height)
	if fontSize != 2 {
		t.Fatalf("%v should fit on a %dx%d board with a 2 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

	// Test when fontSize should equal 1
	listWords = []string{"ThisIsASuperLongWordThatYouCouldNeverDreamOf"}
	width, height = 50, 20

	fontSize = Solve(listWords, width, height)
	if fontSize != 1 {
		t.Fatalf("%v should fit on a %dx%d board with a 1 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

	// Test when fontSize should equal 0
	listWords = []string{"ThisIsASuperLongWordThatYouCouldNeverDreamOf"}
	width, height = 10, 20

	fontSize = Solve(listWords, width, height)
	if fontSize != 0 {
		t.Fatalf("%v should fit on a %dx%d board with a 0 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

	// Test when fontsize < whRatio
	listWords = []string{"SixSix", "SixSix", "SixSix", "SixSix", "SixSix", "SixSix"}
	width, height = 252, 6

	fontSize = Solve(listWords, width, height)
	if fontSize != 6 {
		t.Fatalf("%v should fit on a %dx%d board with a 6 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

	listWords = []string{"hacker", "cup"}
	width, height = 20, 6

	fontSize = Solve(listWords, width, height)
	if fontSize != 3 {
		t.Fatalf("%v should fit on a %dx%d board with a 3 fontSize instead it fits with %d fontSize", listWords, width, height, fontSize)
	}

}
