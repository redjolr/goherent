package experiment_test

import (
	"testing"

	"github.com/redjolr/goherent/cmd/experiment"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert := assert.New(t)
	Test("Passing test", func(t *testing.T) {
		assert.Equal(experiment.Foo(), "Foo")
	}, t)

	Test("Failing test 1", func(t *testing.T) {
		assert.True(false)
	}, t)

	Test("Failing test 2", func(t *testing.T) {
		assert.True(false)
	}, t)

	Test("Passing test 2", func(t *testing.T) {
		assert.Equal(experiment.Foo(), "Foo")
	}, t)

}
