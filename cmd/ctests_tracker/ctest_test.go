package ctests_tracker_test

import (
	"testing"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	. "github.com/redjolr/goherent/pkg"

	"github.com/redjolr/goherent/expect"
)

func TestCtestOutput(t *testing.T) {
	Expect := expect.New(t)

	Test(`
	Given that there is a Ctest with name "someTest" of package "somePackage"
	And that Ctest does not have any output events
	When we call the Output() method on the given ctest
	Then the method will return ""`, func(t *testing.T) {
		// Given
		ctest := ctests_tracker.NewCtest("someTest", "somePackage")

		// When
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("")
	}, t)

	Test(`
	Given that there is a Ctest with name "someTest" of package "somePackage"
	And that Ctest has an output event with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(t *testing.T) {
		// Given
		ctest := ctests_tracker.NewCtest("someTest", "somePackage")

		ctest.RecordOutputEvt(makeCtestOutputEvent("somePackage", "someTest", "some output"))
		// When
		ctestOutput := ctest.Output()
		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given that there is a Ctest with name "someTest" of package "somePackage"
	And that Ctest has 2 output events outputs: "some output 1." and "some output 2."
	When we call the Output() method on the given ctest
	Then the method will return "some output 1.some output 2."`, func(t *testing.T) {
		// Given
		ctest := ctests_tracker.NewCtest("someTest", "somePackage")

		ctest.RecordOutputEvt(makeCtestOutputEvent("somePackage", "someTest", "some output 1."))
		ctest.RecordOutputEvt(makeCtestOutputEvent("somePackage", "someTest", "some output 2."))

		// When
		ctestOutput := ctest.Output()
		// Then
		Expect(ctestOutput).ToEqual("some output 1.some output 2.")
	}, t)

}
