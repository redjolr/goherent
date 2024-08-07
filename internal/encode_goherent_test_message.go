package internal

import (
	"regexp"
	"strings"
)

func EncodeGoherentTestName(testName string) string {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	tabRegex := regexp.MustCompile(`\t`)
	testName = strings.ReplaceAll(testName, " ", ENCODED_WHITESPACE)
	testName = newLineRegex.ReplaceAllString(testName, ENCODED_NEWLINE)
	testName = tabRegex.ReplaceAllString(testName, strings.Repeat(ENCODED_WHITESPACE, 4))
	return testName
}
