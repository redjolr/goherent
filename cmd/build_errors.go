package cmd

import "strings"

// parseBuildErrors groups a test runner's stderr into per-package build error
// text. `go test` writes compiler failures to stderr (not the -json stdout
// stream) as blocks introduced by a "# <import-path>" header line, e.g.:
//
//	# nexus/app/chat/domain/conversation_test [nexus/app/chat/domain/conversation.test]
//	./conversation_test.go:95:8: undefined: conversationNew
//
// The returned map is keyed by the package import path (as it appears in the
// -json "Package" field) and holds the error lines for that package. Output
// before any header is ignored.
func parseBuildErrors(stderr string) map[string]string {
	result := map[string]string{}
	if strings.TrimSpace(stderr) == "" {
		return result
	}

	currentPkg := ""
	var block strings.Builder
	flush := func() {
		if currentPkg != "" && block.Len() > 0 {
			result[currentPkg] += block.String()
		}
		block.Reset()
	}

	// Trim a single trailing newline so the final line doesn't add a spurious blank
	// line to the last package's block.
	for _, line := range strings.Split(strings.TrimSuffix(stderr, "\n"), "\n") {
		if strings.HasPrefix(line, "# ") {
			flush()
			currentPkg = packageFromHeader(line)
			continue
		}
		if currentPkg == "" {
			continue
		}
		block.WriteString(line)
		block.WriteString("\n")
	}
	flush()
	return result
}

// packageFromHeader extracts the package import path from a "# ..." build header.
// The header is either "# <import-path>" or, for a test binary,
// "# <import-path>_test [<import-path>.test]"; in both cases we want the bare
// import path that matches the -json "Package" field.
func packageFromHeader(line string) string {
	s := strings.TrimPrefix(line, "# ")
	if open := strings.IndexByte(s, '['); open != -1 {
		inner := strings.TrimSuffix(s[open+1:], "]")
		inner = strings.TrimSuffix(inner, ".test")
		if inner != "" {
			return inner
		}
	}
	first := s
	if sp := strings.IndexByte(first, ' '); sp != -1 {
		first = first[:sp]
	}
	return strings.TrimSuffix(first, "_test")
}
