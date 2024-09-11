package tests_test

import (
	"errors"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
)

type ErrTypeForTest1 struct{}

func (e ErrTypeForTest1) Error() string {
	return ""
}

type ErrTypeForTest2 struct {
	f int
}

func (e ErrTypeForTest2) Error() string {
	return ""
}

type NonErrType struct {
	f int
}

func (e NonErrType) Error() int {
	return 2
}

func TestToBeError(t *testing.T) {
	t.Run("it should not fail the assertion, if we pass an actual error", func(t *testing.T) {
		assertionErr := assertions.ToBeError(errors.New(""))
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})
	t.Run("it should fail the assertion, if we pass nil", func(t *testing.T) {
		assertionErr := assertions.ToBeError(nil)
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass an integer.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(2)
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a float64.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(2.4)
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a string.", func(t *testing.T) {
		assertionErr := assertions.ToBeError("someString")
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a string.", func(t *testing.T) {
		assertionErr := assertions.ToBeError("someString")
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass true.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(true)
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass false.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(false)
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass an array.", func(t *testing.T) {
		assertionErr := assertions.ToBeError([1]int{8})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a slice.", func(t *testing.T) {
		assertionErr := assertions.ToBeError([]string{"val", "other val"})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a map.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(map[string]string{"k": "v"})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a struct S{f: 2}.", func(t *testing.T) {
		type S struct{ f int }
		assertionErr := assertions.ToBeError(S{f: 2})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should pass the assertion, if we pass a struct ErrTypeForTest1{} with a method Error() string.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(ErrTypeForTest1{})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should pass the assertion, if we pass a struct ErrTypeForTest2{} with a method Error() string.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(ErrTypeForTest2{f: 2})
		if assertionErr != nil {
			t.Errorf("%v", assertionErr)
		}
	})

	t.Run("it should fail the assertion, if we pass a struct NonErrType{f: 2}.", func(t *testing.T) {
		assertionErr := assertions.ToBeError(NonErrType{f: 2})
		if assertionErr == nil {
			t.Errorf("%v", assertionErr)
		}
	})
}
