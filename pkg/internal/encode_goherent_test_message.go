package internal

import (
	"regexp"
	"strings"
)

func EncodeGoherentTestMessage(testMessage string) string {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	tabRegex := regexp.MustCompile(`\t`)
	testMessage = strings.ReplaceAll(testMessage, " ", ENCODED_WHITESPACE)
	testMessage = newLineRegex.ReplaceAllString(testMessage, ENCODED_NEWLINE)
	testMessage = tabRegex.ReplaceAllString(testMessage, ENCODED_TAB)

	return testMessage
}
