package main

func FitsOnBillboard(listWords []string, width, height, fontSize uint) bool {
	if fontSize == 0 {
		return false
	}
	var numberOfCharPerLine int = int(width / fontSize)
	var charSpaceLeftOnCurrentLine int = int(numberOfCharPerLine)
	var numberOfLineLeft int = int(height / fontSize)

	// Font size is too big
	if numberOfLineLeft == 0 || numberOfCharPerLine == 0 {
		return false
	}

	for _, word := range listWords {
		lenCurrentWord := len(word)
		if lenCurrentWord > charSpaceLeftOnCurrentLine {
			numberOfLineLeft--
			charSpaceLeftOnCurrentLine = numberOfCharPerLine
		}
		// Font size too big
		if lenCurrentWord > numberOfCharPerLine || numberOfLineLeft == 0 {
			return false
		}
		charSpaceLeftOnCurrentLine -= lenCurrentWord + 1
	}
	return true
}

func GetRange(listWords []string, width, height, fontSize uint) (lowerBound, upperBound uint) {
	previousSize := fontSize
	if FitsOnBillboard(listWords, width, height, fontSize) {
		fontSize *= 2
		for FitsOnBillboard(listWords, width, height, fontSize) {
			previousSize = fontSize
			fontSize *= 2
		}
		lowerBound = previousSize
		upperBound = fontSize
	} else {
		fontSize /= 2
		for fontSize > 0 && !FitsOnBillboard(listWords, width, height, fontSize) {
			previousSize = fontSize
			fontSize /= 2
		}
		lowerBound = fontSize
		upperBound = previousSize
	}
	return lowerBound, upperBound
}

// Let's avoid having int to float conversion with the math package and just make our own max min functions
func Min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func Max(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}
