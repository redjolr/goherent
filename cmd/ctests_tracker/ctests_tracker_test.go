package ctests_tracker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/expect"

	. "github.com/redjolr/goherent/test"
)

func makeCtestRanEvent(packageName, testName string) events.CtestRanEvent {
	elapsedTime := 1.2
	return events.NewCtestRanEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "run",
			Package: packageName,
			Test:    testName,
			Elapsed: &elapsedTime,
		},
	)
}

func makeCtestPassedEvent(packageName, testName string) events.CtestPassedEvent {
	timeElapsed := 1.2
	return events.NewCtestPassedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "pass",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func makeCtestFailedEvent(packageName, testName string) events.CtestFailedEvent {
	timeElapsed := 1.2
	return events.NewCtestFailedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "pass",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func makeCtestOutputEvent(packageName, testName, output string) events.CtestOutputEvent {
	return events.NewCtestOutputEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "output",
			Test:    testName,
			Package: packageName,
			Output:  output,
		},
	)
}

func makeCtestSkippedEvent(packageName, testName string) events.CtestSkippedEvent {
	timeElapsed := 1.2
	return events.NewCtestSkippedEvent(
		events.JsonTestEvent{
			Time:    time.Now(),
			Action:  "skip",
			Test:    testName,
			Package: packageName,
			Elapsed: &timeElapsed,
		},
	)
}

func TestHandleCtestRanEvent(t *testing.T) {
	Test(`
	Given an empty CtestsTracker
	When a CtestRanEvent for a test named "someTest" in package "somePackage"
	Then the CtestsTracker should contain that PackageUnderTest
	And a running Ctest with that name should be added. 
	`, func(Expect expect.F) {
		tracker := ctests_tracker.NewCtestsTracker()

		ctestRanEvent := makeCtestRanEvent("somePackage", "someTest")

		tracker.HandleCtestRanEvent(ctestRanEvent)
		ctest := tracker.FindCtestWithNameInPackage("someTest", "somePackage")

		Expect(tracker.ContainsPackageUtWithName("somePackage")).ToEqual(true)

		Expect(ctest.Name()).ToEqual("someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.IsRunning()).ToEqual(true)
	}, t)
}

func TestHandleCtestPassedEvent(t *testing.T) {
	Test(`
	Given that there is an empty CtestsTracker
	When a CtestPassedEvent for test "someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "someTest")

		tracker.HandleCtestPassedEvent(ctestPassedEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("someTest", "somePackage")

		// Then
		Expect(ctest.Name()).ToEqual("someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasPassed()).ToEqual(true)
	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "someTest" in package "somePackage"
	When a CtestPassedEvent for test "someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "someTest", "some output")
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "someTest")
		tracker.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		tracker.HandleCtestPassedEvent(ctestPassedEvt)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("someTest", "somePackage")

		Expect(ctest.Name()).ToEqual("someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasPassed()).ToEqual(true)

	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "someTest" in package "somePackage"
	And a CtestPassedEvent for test "someTest" from package "somePackage" occurrs
	When a second CtestPassedEvent for test "someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestPassedEvt1 := makeCtestPassedEvent("somePackage", "someTest")
		ctestPassedEvt2 := makeCtestPassedEvent("somePackage", "someTest")

		tracker.HandleCtestPassedEvent(ctestPassedEvt1)

		// When
		tracker.HandleCtestPassedEvent(ctestPassedEvt2)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("someTest", "somePackage")
		Expect(ctest.Name()).ToEqual("someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasPassed()).ToEqual(true)
		Expect(
			tracker.FindPackageWithName("somePackage").CtestsCount(),
		).ToEqual(1)
	}, t)
}

func TestHandleCtestFailedEvent(t *testing.T) {
	Test(`
	Given that there is an empty CtestsTracker
	When a CtestFailedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	Then a failed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "ParentTest/someTest")

		tracker.HandleCtestFailedEvent(ctestFailedEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")

		// Then
		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasFailed()).ToEqual(true)
	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "ParentTest/someTest" in package "somePackage"
	When a CtestFailedEvent for test "someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")
		ctestFailedEvt := makeCtestFailedEvent("somePackage", "ParentTest/someTest")
		tracker.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		tracker.HandleCtestFailedEvent(ctestFailedEvt)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")

		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasFailed()).ToEqual(true)

	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "ParentTest/someTest" in package "somePackage"
	And a CtestFailedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	When a second CtestPassedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestFailedEvt1 := makeCtestFailedEvent("somePackage", "ParentTest/someTest")
		ctestFailedEvt2 := makeCtestFailedEvent("somePackage", "ParentTest/someTest")

		tracker.HandleCtestFailedEvent(ctestFailedEvt1)

		// When
		tracker.HandleCtestFailedEvent(ctestFailedEvt2)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.HasFailed()).ToEqual(true)
		Expect(
			tracker.FindPackageWithName("somePackage").CtestsCount(),
		).ToEqual(1)
	}, t)
}

func TestHandleCtestSkippedEvent(t *testing.T) {
	Test(`
	Given that there is an empty CtestsTracker
	When a CtestSkippedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	Then a failed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "ParentTest/someTest")

		tracker.HandleCtestSkippedEvent(ctestSkippedEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")

		// Then
		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.IsSkipped()).ToEqual(true)
	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "ParentTest/someTest" in package "somePackage"
	When a CtestSkippedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")
		ctestSkippedEvt := makeCtestSkippedEvent("somePackage", "ParentTest/someTest")
		tracker.HandleCtestOutputEvent(ctestOutputEvt)
		// When
		tracker.HandleCtestSkippedEvent(ctestSkippedEvt)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")

		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.IsSkipped()).ToEqual(true)

	}, t)

	Test(`
	Given that there is a CtestsTracker
	And an CtestOutputEvent has occurred for test "ParentTest/someTest" in package "somePackage"
	And a CtestSkippedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	When a second CtestPassedEvent for test "ParentTest/someTest" from package "somePackage" occurrs
	Then a passed Ctest will be stored in the CtestsTracker`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctestSkippedEvt1 := makeCtestSkippedEvent("somePackage", "ParentTest/someTest")
		ctestSkippedEvt2 := makeCtestSkippedEvent("somePackage", "ParentTest/someTest")

		tracker.HandleCtestSkippedEvent(ctestSkippedEvt1)

		// When
		tracker.HandleCtestSkippedEvent(ctestSkippedEvt2)

		// Then
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		Expect(ctest.Name()).ToEqual("ParentTest/someTest")
		Expect(ctest.PackageName()).ToEqual("somePackage")
		Expect(ctest.IsSkipped()).ToEqual(true)
		Expect(
			tracker.FindPackageWithName("somePackage").CtestsCount(),
		).ToEqual(1)
	}, t)
}

func TestNewCtestOutput(t *testing.T) {
	Test(`
	Given that there is a Ctest with name "ParentTest/someTest" of package "somePackage"
	And a CtestOutputEvent has occurred with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertCtest(ctests_tracker.NewCtest("ParentTest/someTest", "somePackage"))
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")

		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given that there is a running Ctest with name "someTest" of package "somePackage"
	And a CtestOutputEvent has occurred with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		ctestRanEvt := makeCtestRanEvent("somePackage", "ParentTest/someTest")
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")

		tracker := ctests_tracker.NewCtestsTracker()
		tracker.HandleCtestRanEvent(ctestRanEvt)
		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given that there is a passed Ctest with name "ParentTest/someTest" of package "somePackage"
	And a CtestOutputEvent has occurred with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		ctestPassedEvt := makeCtestPassedEvent("somePackage", "ParentTest/someTest")
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")

		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertCtest(ctests_tracker.NewPassedCtest(ctestPassedEvt))
		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given that there is a failed Ctest with name "ParentTest/someTest" of package "somePackage"
	And a CtestOutputEvent has occurred with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		ctestPassedEvt := makeCtestFailedEvent("somePackage", "ParentTest/someTest")
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")

		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertCtest(ctests_tracker.NewFailedCtest(ctestPassedEvt))
		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given that there is a package "somePackage"
	And a CtestOutputEvent has occurred for "ParentTest/someTest" of "somePackage" with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")

		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")
		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)

	Test(`
	Given a CtestOutputEvent has occurred for "ParentTest/someTest" of "somePackage" with output "some output"
	When we call the Output() method on the given ctest
	Then the method will return "some output"`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		ctestOutputEvt := makeCtestOutputEvent("somePackage", "ParentTest/someTest", "some output")
		tracker.HandleCtestOutputEvent(ctestOutputEvt)

		// When
		ctest := tracker.FindCtestWithNameInPackage("ParentTest/someTest", "somePackage")
		ctestOutput := ctest.Output()

		// Then
		Expect(ctestOutput).ToEqual("some output")
	}, t)
}

func TestInsertCtest(t *testing.T) {
	Test(`
	Given that we have an empty TestsTracker
	When we call the InsertCtest() method with a Ctest as an argument
	The Ctest will be added and returned from the method
	And a package with the name of the Ctest's packageName will be added to the tracker
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		// When
		ctest := ctests_tracker.NewCtest("testName", "packageName")
		ctestReturned := tracker.InsertCtest(ctest)

		// Then
		Expect(ctestReturned).ToEqual(ctest)
		Expect(tracker.ContainsPackageUtWithName("packageName")).ToBeTrue()
		testInPackage := tracker.FindCtestWithNameInPackage("testName", "packageName")
		Expect(testInPackage).NotToBeNil()
		fmt.Println("\n\n\n YAYYYYY")
	}, t)

	Test(`
	Given that we have a CtestsTracker
	And that tracker has a PackageUnderTest with name "packageName"
	And that PackageUnderTest has a Ctest with name "ctestName1"
	When we call the InsertCtest() method with a Ctest { name: "ctestName2", packageName: "packageName" }
	Then the Ctest will be added to the existing PackageUnderTest
	And no existing packageUnderTest will be added to the tracker
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageName")
		tracker.InsertCtest(ctest1)
		Expect(tracker.PackagesCount()).ToEqual(1)

		// When
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageName")
		tracker.InsertCtest(ctest2)

		// Then
		Expect(tracker.PackagesCount()).ToEqual(1)
		testInPackage := tracker.FindCtestWithNameInPackage("ctestName2", "packageName")
		Expect(testInPackage).NotToBeNil()
	}, t)

	Test(`
	Given that we have a CtestsTracker
	And that tracker has a PackageUnderTest with name "packageName1"
	And that PackageUnderTest has a Ctest with name "ctestName1"
	When we call the InsertCtest() method with a Ctest { name: "ctestName2", packageName: "packageName2" }
	Then a new PackageUnderTest will be created
	And the ctestName2 Ctest will be added to that new package
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageName1")
		tracker.InsertCtest(ctest1)
		Expect(tracker.PackagesCount()).ToEqual(1)

		// When
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageName2")
		tracker.InsertCtest(ctest2)

		// Then
		testInPackage := tracker.FindCtestWithNameInPackage("ctestName2", "packageName2")
		Expect(testInPackage).NotToBeNil()
		Expect(tracker.PackagesCount()).ToEqual(2)
	}, t)
}

func TestIsCtestFirstOfItsPackage(t *testing.T) {
	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has only one Ctest with name "ctestName"
	When we check if that ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return true
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewCtest("testName", "packageUtName")
		tracker.InsertCtest(ctest)

		// When
		isFirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest)

		//Then
		Expect(isFirstInPackage).ToBeTrue()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has two Ctests with names: "ctestName1" and "ctestName2"
	When we check if "ctestName1" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return true
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageUtName")
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageUtName")

		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		isCtest1FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest1)

		//Then
		Expect(isCtest1FirstInPackage).ToBeTrue()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has two Ctests with names: "ctestName1" and "ctestName2"
	When we check if "ctestName2" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageUtName")
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageUtName")

		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		isCtest2FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest2)

		//Then
		Expect(isCtest2FirstInPackage).ToBeFalse()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has 3 Ctests with names: "ctestName1", "ctestName2", "ctestName3"
	When we check if "ctestName2" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageUtName")
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageUtName")
		ctest3 := ctests_tracker.NewCtest("ctestName3", "packageUtName")

		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)
		tracker.InsertCtest(ctest3)

		// When
		isCtest2FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest2)

		//Then
		Expect(isCtest2FirstInPackage).ToBeFalse()
	}, t)

	Test(`
	Given that we have an empty CtestTracker
	When we check if "ctestName" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		ctest := ctests_tracker.NewCtest("ctestName", "packageUtName")
		isCtestFirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest)

		//Then
		Expect(isCtestFirstInPackage).ToBeFalse()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 ("someTestName")
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has q Ctests: ctest1InPackage2 ("ctest1InPackage2")
	When we check if a Ctest {name: "someTestName", packageName: "packageUtName2" } is the first of its package
	Then the method should return false
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1InPackage1 := ctests_tracker.NewCtest("someTestName", "packageUtName1")
		ctest1InPackage2 := ctests_tracker.NewCtest("ctest1InPackage2", "packageUtName2")

		tracker.InsertCtest(ctest1InPackage1)
		tracker.InsertCtest(ctest1InPackage2)

		// When
		someTestNamePackage2 := ctests_tracker.NewCtest("someTestName", "packageUtName2")
		isSomeTestNameFirstOfPackage2 := tracker.IsCtestFirstOfItsPackage(someTestNamePackage2)

		//Then
		Expect(isSomeTestNameFirstOfPackage2).ToBeFalse()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 with name: "someTestName"
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has 2 Ctests: ctest1InPackage2 ("ctest1InPackage2"), ctest2InPackage2 ("someTestName")
	When we check if ctest2InPackage2 is the first of its package
	Then the method should return false
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1InPackage1 := ctests_tracker.NewCtest("someTestName", "packageUtName1")
		ctest1InPackage2 := ctests_tracker.NewCtest("ctest1InPackage2", "packageUtName2")
		ctest2InPackage2 := ctests_tracker.NewCtest("someTestName", "packageUtName2")

		tracker.InsertCtest(ctest1InPackage1)
		tracker.InsertCtest(ctest1InPackage2)
		tracker.InsertCtest(ctest2InPackage2)

		// When
		isCtest2FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest2InPackage2)

		//Then
		Expect(isCtest2FirstInPackage).ToBeFalse()
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 with name: "someTestName"
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has 2 Ctests: ctest1InPackage2 ("ctest1InPackage2"), ctest2InPackage2 ("someTestName")
	When we check if ctest1InPackage1 is the first of its package
	Then the method should return true
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1InPackage1 := ctests_tracker.NewCtest("someTestName", "packageUtName1")
		ctest1InPackage2 := ctests_tracker.NewCtest("ctest1InPackage2", "packageUtName2")
		ctest2InPackage2 := ctests_tracker.NewCtest("someTestName", "packageUtName2")

		tracker.InsertCtest(ctest1InPackage1)
		tracker.InsertCtest(ctest1InPackage2)
		tracker.InsertCtest(ctest2InPackage2)

		// When
		isCtest1FirstInPackage1 := tracker.IsCtestFirstOfItsPackage(ctest1InPackage1)

		//Then
		Expect(isCtest1FirstInPackage1).ToBeTrue()
	}, t)
}

func TestRunningTestsCount(t *testing.T) {
	Test(`
	Given that the CtestTracker does not have any Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(0)
	}, t)

	Test(`
	Given that the CtestTracker has a passed Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewPassedCtest(makeCtestPassedEvent("somePackage", "testName"))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(0)
	}, t)

	Test(`
	Given that the CtestTracker has a failed Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewFailedCtest(makeCtestFailedEvent("somePackage", "testName"))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(0)
	}, t)

	Test(`
	Given that the CtestTracker has a running Ctest in it
	When we execute the RunningCtestsCount()
	Then the return value should be 1
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage", "testName"))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(1)
	}, t)

	Test(`
	Given that the CtestTracker has two running Ctests in the same package named "somePackage"
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage", "testName"))

		ctest2 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage", "testName2"))
		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(2)
	}, t)

	Test(`
	Given that the CtestTracker has two running Ctests in different packages
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage", "testName"))

		ctest2 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage2", "testName2"))
		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(2)
	}, t)

	Test(`
	Given that the CtestTracker has 1 passed test and a running test in package "somePackage"
	And 1 running test in package "somePackage2"
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		runningCtest1 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage", "testName1"))

		passingCtest := ctests_tracker.NewPassedCtest(makeCtestPassedEvent("somePackage", "testName2"))

		runningCtest2 := ctests_tracker.NewRunningCtest(makeCtestRanEvent("somePackage2", "testName3"))
		tracker.InsertCtest(runningCtest1)
		tracker.InsertCtest(passingCtest)
		tracker.InsertCtest(runningCtest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		Expect(runningCtestsCnt).ToEqual(2)
	}, t)
}

func TestDeletePackage(t *testing.T) {
	Test(`
	Given that there is a CtestTracker with no packages,
	When we try to delete a random package
	Then nothing will happen and the CtestsTracker will still have no packages.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		randomPackage := ctests_tracker.NewPackageUnderTest("somePackage")
		tracker.DeletePackage(&randomPackage)

		// Then
		Expect(tracker.PackagesCount()).ToEqual(0)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package named "somePackage",
	When we try to delete that package
	Then the package will be deleted and the tracker will have 0 packages.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		somePackage := tracker.FindPackageWithName("somePackage")
		tracker.DeletePackage(somePackage)

		// Then
		Expect(tracker.PackagesCount()).ToEqual(0)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package,
	When we try to delete another package 
	Then nothing will happen and the CtestsTracker will have 1 package.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		someOtherPackage := ctests_tracker.NewPackageUnderTest("someOtherPackage")
		tracker.DeletePackage(&someOtherPackage)

		// Then
		Expect(tracker.PackagesCount()).ToEqual(1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package with name "somePackage",
	When we try to delete another package which also has that name (but it is a different instance)
	Then nothing will happen and the CtestsTracker will have 1 packages.`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		someOtherPackage := ctests_tracker.NewPackageUnderTest("somePackage")
		tracker.DeletePackage(&someOtherPackage)

		// Then
		Expect(tracker.PackagesCount()).ToEqual(1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 2 package with names "somePackage 1" and "somePackage 2",
	When we try to delete the "somePackage 1" package
	Then the "somePackage 1" package will be deleted`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")

		// When
		somePackage1 := tracker.FindPackageWithName("somePackage 1")
		tracker.DeletePackage(somePackage1)

		// Then
		Expect(tracker.Packages()[0].Name()).ToEqual("somePackage 2")
		Expect(tracker.PackagesCount()).ToEqual(1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 2 package with names "somePackage 1" and "somePackage 2",
	When we try to delete the "somePackage 2" package
	Then the "somePackage 2" package will be deleted`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")

		// When
		somePackage2 := tracker.FindPackageWithName("somePackage 2")
		tracker.DeletePackage(somePackage2)

		// Then
		Expect(tracker.Packages()[0].Name()).ToEqual("somePackage 1")
		Expect(tracker.PackagesCount()).ToEqual(1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 1" package
	Then the "somePackage 1" package will be deleted`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage1 := tracker.FindPackageWithName("somePackage 1")
		tracker.DeletePackage(somePackage1)

		// Then
		Expect(tracker.Packages()[0].Name()).ToEqual("somePackage 2")
		Expect(tracker.Packages()[1].Name()).ToEqual("somePackage 3")
		Expect(tracker.PackagesCount()).ToEqual(2)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 2" package
	Then the "somePackage 2" package will be deleted`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage2 := tracker.FindPackageWithName("somePackage 2")
		tracker.DeletePackage(somePackage2)

		// Then
		Expect(tracker.Packages()[0].Name()).ToEqual("somePackage 1")
		Expect(tracker.Packages()[1].Name()).ToEqual("somePackage 3")
		Expect(tracker.PackagesCount()).ToEqual(2)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 3" package
	Then the "somePackage 2" package will be deleted`, func(Expect expect.F) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage3 := tracker.FindPackageWithName("somePackage 3")
		tracker.DeletePackage(somePackage3)

		// Then
		Expect(tracker.Packages()[0].Name()).ToEqual("somePackage 1")
		Expect(tracker.Packages()[1].Name()).ToEqual("somePackage 2")
		Expect(tracker.PackagesCount()).ToEqual(2)
	}, t)
}
