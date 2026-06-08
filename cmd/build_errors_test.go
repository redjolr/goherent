package cmd

import "testing"

func TestParseBuildErrors(t *testing.T) {
	t.Run("maps a test-binary build failure to its import path", func(t *testing.T) {
		stderr := "# nexus/app/chat/domain/conversation_test [nexus/app/chat/domain/conversation.test]\n" +
			"./conversation_test.go:95:8: undefined: conversationNew\n"

		got := parseBuildErrors(stderr)

		want := "./conversation_test.go:95:8: undefined: conversationNew\n"
		if got["nexus/app/chat/domain/conversation"] != want {
			t.Fatalf("got %q for package, want %q", got["nexus/app/chat/domain/conversation"], want)
		}
		if len(got) != 1 {
			t.Fatalf("expected exactly one package, got %d: %v", len(got), got)
		}
	})

	t.Run("maps a plain (non-test) build failure header", func(t *testing.T) {
		stderr := "# nexus/app/chat/domain/conversation\n" +
			"./conversation.go:10:2: syntax error\n"

		got := parseBuildErrors(stderr)

		if got["nexus/app/chat/domain/conversation"] != "./conversation.go:10:2: syntax error\n" {
			t.Fatalf("unexpected mapping: %v", got)
		}
	})

	t.Run("separates errors from multiple packages", func(t *testing.T) {
		stderr := "# pkg/a [pkg/a.test]\n./a_test.go:1:1: undefined: a\n" +
			"# pkg/b [pkg/b.test]\n./b_test.go:2:2: undefined: b\n"

		got := parseBuildErrors(stderr)

		if got["pkg/a"] != "./a_test.go:1:1: undefined: a\n" {
			t.Fatalf("pkg/a: got %q", got["pkg/a"])
		}
		if got["pkg/b"] != "./b_test.go:2:2: undefined: b\n" {
			t.Fatalf("pkg/b: got %q", got["pkg/b"])
		}
	})

	t.Run("returns empty for blank stderr", func(t *testing.T) {
		if got := parseBuildErrors("   \n"); len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})
}
