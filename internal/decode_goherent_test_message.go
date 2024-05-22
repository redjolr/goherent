package internal

import (
	"strings"
)

func DecodeGoherentTestMessage(encodedTestName string) string {
	decoded := strings.ReplaceAll(encodedTestName, ENCODED_WHITESPACE, " ")
	decoded = strings.ReplaceAll(decoded, ENCODED_NEWLINE, "\n")
	decoded = strings.ReplaceAll(decoded, ENCODED_TAB, "\t")
	return decoded
}
