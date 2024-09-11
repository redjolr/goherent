package tests_test

import (
	"errors"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

type ErrorTypeForTest1 struct{}

func (e ErrorTypeForTest1) Error() string {
	return ""
}

type ErrorTypeForTest2 struct {
	f int
}

func (e ErrorTypeForTest2) Error() string {
	return ""
}

type NonErrorType struct {
	f int
}

func (e NonErrorType) Error() int {
	return 2
}

func TestNotToBeError(t *testing.T) {
	t.Run("it should not fail the assertion, if we pass an actual error", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(errors.New(""))
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})
	t.Run("it should not fail the assertion, if we pass nil", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(nil)
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass an integer.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(2)
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a float64.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(2.4)
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a string.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError("someString")
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a string.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError("someString")
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass true.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(true)
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass false.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(false)
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass an array.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError([1]int{8})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a slice.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError([]string{"val", "other val"})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a map.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(map[string]string{"k": "v"})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a struct S{f: 2}.", func(t *testing.T) {
		type S struct{ f int }
		assertionErr := assertions.NotToBeError(S{f: 2})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a struct ErrorTypeForTest1{} with a method Error() string.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(ErrorTypeForTest1{})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a struct ErrorTypeForTest2{} with a method Error() string.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(ErrorTypeForTest2{f: 2})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should not fail the assertion, if we pass a struct NonErrorType{f: 2}.", func(t *testing.T) {
		assertionErr := assertions.NotToBeError(NonErrorType{f: 2})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})
}
