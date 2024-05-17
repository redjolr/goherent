package main

import (
	"regexp"
	"testing"

	"github.com/redjolr/goherent/internal"
)

func Test(message string, fn func(t *testing.T), t *testing.T) {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	whitespaceRegex := regexp.MustCompile(`\s`)
	message = newLineRegex.ReplaceAllString(message, internal.NewLineMessageSeparator)
	message = whitespaceRegex.ReplaceAllString(message, internal.WhitespaceMessageSeparator)
	t.Run(message, fn)
}
