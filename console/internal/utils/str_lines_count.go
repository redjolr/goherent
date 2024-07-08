package utils

func StrLinesCount(str string) int {
	lines := SplitStringByNewLine(str)
	return len(lines)
}
