package utils

import "regexp"

func SplitStringByNewLine(strToSplit string) []string {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	lines := newLineRegex.Split(strToSplit, -1)
	return lines
}
