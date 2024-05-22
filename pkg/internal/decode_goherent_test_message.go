package internal

import (
	"fmt"
	"strings"
)

func DecodeGoherentTestMessage(encodedTestMessage string) string {
	// newLineRegex := regexp.MustCompile(`\r?\n`)
	// tabRegex := regexp.MustCompile(`\t`)
	// testMessage = strings.ReplaceAll(testMessage, " ", ENCODED_WHITESPACE)
	// testMessage = newLineRegex.ReplaceAllString(testMessage, ENCODED_NEWLINE)
	// testMessage = tabRegex.ReplaceAllString(testMessage, ENCODED_TAB)
	decoded := strings.ReplaceAll(encodedTestMessage, ENCODED_WHITESPACE, " ")
	decoded = strings.ReplaceAll(decoded, ENCODED_NEWLINE, "\n")
	decoded = strings.ReplaceAll(decoded, ENCODED_TAB, "\t")
	fmt.Println("\n\n\n DECODED", decoded)
	return decoded
}
