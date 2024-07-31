package foo2_test

import (
	"testing"

	. "github.com/redjolr/goherent/pkg"
)

func TestFoo(t *testing.T) {
	Test("Passing test", func(t *testing.T) {
	}, t)

	Test("This is the first Failing test.", func(t *testing.T) {
		t.Error("Failing test 1")
	}, t)

	Test("This test is failing too", func(t *testing.T) {
		t.Error("Failing test 1")
	}, t)

	Test(`This is a passing test. If you dont believe me,
     check it for yourself.`, func(t *testing.T) {
	}, t)

	Test("Passing test 2", func(t *testing.T) {
		t.Skip()
	}, t)

	Test("This is the first Failing test.", func(t *testing.T) {
		t.Error("Failing test 1")
	}, t)

	Test("This test is failing too", func(t *testing.T) {
		t.Error("Failing test 1")
	}, t)

	Test(`This is a passing test. If you dont believe me,
     check it for yourself.`, func(t *testing.T) {
		t.Skip()
	}, t)

	Test("Passing test 2", func(t *testing.T) {
	}, t)

}
