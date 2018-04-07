package main

func FitsOnBillboard(listWords []string, width, height, fontSize uint) bool {
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
