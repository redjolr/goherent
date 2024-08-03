package ctests_tracker_test

import (
	"testing"
	"time"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
	. "github.com/redjolr/goherent/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewCtestRanEvent(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given an empty CtestsTracker
	When a CtestRanEvent for a certain test in a certain package is received
	Then the CtestsTracker should contain that PackageUnderTest
	`, func(t *testing.T) {
		tracker := ctests_tracker.NewCtestsTracker()

		packageName := "somePackageName"
		ctestRanEvent := ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Package: packageName,
				Test:    "someTestName",
			},
		)

		tracker.NewCtestRanEvent(ctestRanEvent)
		assert.True(tracker.ContainsPackageUtWithName(packageName))
	}, t)
}

func TestInsertCtest(t *testing.T) {
	assert := assert.New(t)
	Test(`
	Given that we have an empty TestsTracker
	When we call the InsertCtest() method with a Ctest as an argument
	The Ctest will be added and returned from the method
	And a package with the name of the Ctest's packageName will be added to the tracker
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		// When
		ctest := ctests_tracker.NewCtest("testName", "packageName")
		ctestReturned := tracker.InsertCtest(ctest)

		// Then
		assert.Equal(ctestReturned, ctest)
		assert.True(tracker.ContainsPackageUtWithName("packageName"))
		testInPackage := tracker.FindCtestWithNameInPackage("testName", "packageName")
		assert.NotNil(testInPackage)
	}, t)

	Test(`
	Given that we have a CtestsTracker
	And that tracker has a PackageUnderTest with name "packageName"
	And that PackageUnderTest has a Ctest with name "ctestName1"
	When we call the InsertCtest() method with a Ctest { name: "ctestName2", packageName: "packageName" }
	Then the Ctest will be added to the existing PackageUnderTest
	And no existing packageUnderTest will be added to the tracker
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageName")
		tracker.InsertCtest(ctest1)
		assert.Equal(tracker.PackagesCount(), 1)

		// When
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageName")
		tracker.InsertCtest(ctest2)

		// Then
		assert.Equal(tracker.PackagesCount(), 1)
		testInPackage := tracker.FindCtestWithNameInPackage("ctestName2", "packageName")
		assert.NotNil(testInPackage)
	}, t)

	Test(`
	Given that we have a CtestsTracker
	And that tracker has a PackageUnderTest with name "packageName1"
	And that PackageUnderTest has a Ctest with name "ctestName1"
	When we call the InsertCtest() method with a Ctest { name: "ctestName2", packageName: "packageName2" }
	Then a new PackageUnderTest will be created
	And the ctestName2 Ctest will be added to that new package
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageName1")
		tracker.InsertCtest(ctest1)
		assert.Equal(tracker.PackagesCount(), 1)

		// When
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageName2")
		tracker.InsertCtest(ctest2)

		// Then
		testInPackage := tracker.FindCtestWithNameInPackage("ctestName2", "packageName2")
		assert.NotNil(testInPackage)
		assert.Equal(tracker.PackagesCount(), 2)
	}, t)
}

func TestIsCtestFirstOfItsPackage(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has only one Ctest with name "ctestName"
	When we check if that ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return true
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewCtest("testName", "packageUtName")
		tracker.InsertCtest(ctest)

		// When
		isFirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest)

		//Then
		assert.True(isFirstInPackage)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has two Ctests with names: "ctestName1" and "ctestName2"
	When we check if "ctestName1" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return true
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageUtName")
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageUtName")

		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		isCtest1FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest1)

		//Then
		assert.True(isCtest1FirstInPackage)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has two Ctests with names: "ctestName1" and "ctestName2"
	When we check if "ctestName2" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewCtest("ctestName1", "packageUtName")
		ctest2 := ctests_tracker.NewCtest("ctestName2", "packageUtName")

		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		isCtest2FirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest2)

		//Then
		assert.False(isCtest2FirstInPackage)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName"
	And the given PackageUnderTest has 3 Ctests with names: "ctestName1", "ctestName2", "ctestName3"
	When we check if "ctestName2" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(t *testing.T) {
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
		assert.False(isCtest2FirstInPackage)
	}, t)

	Test(`
	Given that we have an empty CtestTracker
	When we check if "ctestName" Ctest is the first in the package with the IsCtestFirstOfItsPackage() method
	Then the method should return false
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		ctest := ctests_tracker.NewCtest("ctestName", "packageUtName")
		isCtestFirstInPackage := tracker.IsCtestFirstOfItsPackage(ctest)

		//Then
		assert.False(isCtestFirstInPackage)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 ("someTestName")
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has q Ctests: ctest1InPackage2 ("ctest1InPackage2")
	When we check if a Ctest {name: "someTestName", packageName: "packageUtName2" } is the first of its package
	Then the method should return false
	`, func(t *testing.T) {
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
		assert.False(isSomeTestNameFirstOfPackage2)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 with name: "someTestName"
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has 2 Ctests: ctest1InPackage2 ("ctest1InPackage2"), ctest2InPackage2 ("someTestName")
	When we check if ctest2InPackage2 is the first of its package
	Then the method should return false
	`, func(t *testing.T) {
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
		assert.False(isCtest2FirstInPackage)
	}, t)

	Test(`
	Given that we have a CtestTracker
	And that CtestTracker has a PackageUnderTest with name "packageUtName1"
	And the "packageUtName1" PackageUnderTest has 1 Ctests ctest1InPackage1 with name: "someTestName"
	And that CtestTracker has a PackageUnderTest with name "packageUtName2"
	And the "packageUtName2" has 2 Ctests: ctest1InPackage2 ("ctest1InPackage2"), ctest2InPackage2 ("someTestName")
	When we check if ctest1InPackage1 is the first of its package
	Then the method should return true
	`, func(t *testing.T) {
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
		assert.True(isCtest1FirstInPackage1)
	}, t)
}

func TestRunningTestsCount(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that the CtestTracker does not have any Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0
	`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 0)
	}, t)

	Test(`
	Given that the CtestTracker has a passed Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewPassedCtest(ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 0)
	}, t)

	Test(`
	Given that the CtestTracker has a failed Ctests in it
	When we execute the RunningCtestsCount()
	Then the return value should be 0
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewFailedCtest(ctest_failed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 0)
	}, t)

	Test(`
	Given that the CtestTracker has a running Ctest in it
	When we execute the RunningCtestsCount()
	Then the return value should be 1
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		ctest := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(ctest)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 1)
	}, t)

	Test(`
	Given that the CtestTracker has two running Ctests in the same package named "somePackage"
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		))

		ctest2 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName2",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 2)
	}, t)

	Test(`
	Given that the CtestTracker has two running Ctests in different packages
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		ctest1 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName",
				Elapsed: &elapsedTime,
			},
		))

		ctest2 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage2",
				Test:    "testName2",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(ctest1)
		tracker.InsertCtest(ctest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 2)
	}, t)

	Test(`
	Given that the CtestTracker has 1 passed test and a running test in package "somePackage"
	And 1 running test in package "somePackage2"
	When we execute the RunningCtestsCount()
	Then the return value should be 2
	`, func(t *testing.T) {
		// Given
		elapsedTime := 2.3
		tracker := ctests_tracker.NewCtestsTracker()
		runningCtest1 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName1",
				Elapsed: &elapsedTime,
			},
		))

		passingCtest := ctests_tracker.NewPassedCtest(ctest_passed_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage",
				Test:    "testName2",
				Elapsed: &elapsedTime,
			},
		))

		runningCtest2 := ctests_tracker.NewRunningCtest(ctest_ran_event.NewFromJsonTestEvent(
			events.JsonTestEvent{
				Time:    time.Now(),
				Action:  "pass",
				Package: "somePackage2",
				Test:    "testName3",
				Elapsed: &elapsedTime,
			},
		))
		tracker.InsertCtest(runningCtest1)
		tracker.InsertCtest(passingCtest)
		tracker.InsertCtest(runningCtest2)

		// When
		runningCtestsCnt := tracker.RunningCtestsCount()

		// Then
		assert.Equal(runningCtestsCnt, 2)
	}, t)
}

func TestDeletePackage(t *testing.T) {
	assert := assert.New(t)

	Test(`
	Given that there is a CtestTracker with no packages,
	When we try to delete a random package
	Then nothing will happen and the CtestsTracker will still have no packages.`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()

		// When
		randomPackage := ctests_tracker.NewPackageUnderTest("somePackage")
		tracker.DeletePackage(&randomPackage)

		// Then
		assert.Equal(tracker.PackagesCount(), 0)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package named "somePackage",
	When we try to delete that package
	Then the package will be deleted and the tracker will have 0 packages.`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		somePackage := tracker.FindPackageWithName("somePackage")
		tracker.DeletePackage(somePackage)

		// Then
		assert.Equal(tracker.PackagesCount(), 0)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package,
	When we try to delete another package 
	Then nothing will happen and the CtestsTracker will have 1 package.`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		someOtherPackage := ctests_tracker.NewPackageUnderTest("someOtherPackage")
		tracker.DeletePackage(&someOtherPackage)

		// Then
		assert.Equal(tracker.PackagesCount(), 1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 1 package with name "somePackage",
	When we try to delete another package which also has that name (but it is a different instance)
	Then nothing will happen and the CtestsTracker will have 1 packages.`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage")

		// When
		someOtherPackage := ctests_tracker.NewPackageUnderTest("somePackage")
		tracker.DeletePackage(&someOtherPackage)

		// Then
		assert.Equal(tracker.PackagesCount(), 1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 2 package with names "somePackage 1" and "somePackage 2",
	When we try to delete the "somePackage 1" package
	Then the "somePackage 1" package will be deleted`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")

		// When
		somePackage1 := tracker.FindPackageWithName("somePackage 1")
		tracker.DeletePackage(somePackage1)

		// Then
		assert.Equal(
			tracker.Packages()[0].Name(),
			"somePackage 2",
		)
		assert.Equal(tracker.PackagesCount(), 1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 2 package with names "somePackage 1" and "somePackage 2",
	When we try to delete the "somePackage 2" package
	Then the "somePackage 2" package will be deleted`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")

		// When
		somePackage2 := tracker.FindPackageWithName("somePackage 2")
		tracker.DeletePackage(somePackage2)

		// Then
		assert.Equal(
			tracker.Packages()[0].Name(),
			"somePackage 1",
		)
		assert.Equal(tracker.PackagesCount(), 1)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 1" package
	Then the "somePackage 1" package will be deleted`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage1 := tracker.FindPackageWithName("somePackage 1")
		tracker.DeletePackage(somePackage1)

		// Then
		assert.Equal(
			tracker.Packages()[0].Name(),
			"somePackage 2",
		)
		assert.Equal(
			tracker.Packages()[1].Name(),
			"somePackage 3",
		)
		assert.Equal(tracker.PackagesCount(), 2)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 2" package
	Then the "somePackage 2" package will be deleted`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage2 := tracker.FindPackageWithName("somePackage 2")
		tracker.DeletePackage(somePackage2)

		// Then
		assert.Equal(
			tracker.Packages()[0].Name(),
			"somePackage 1",
		)
		assert.Equal(
			tracker.Packages()[1].Name(),
			"somePackage 3",
		)
		assert.Equal(tracker.PackagesCount(), 2)
	}, t)

	Test(`
	Given that there is a CtestTracker with 3 packages with names "somePackage 1", "somePackage 2", "somePackage 3",
	When we try to delete the "somePackage 3" package
	Then the "somePackage 2" package will be deleted`, func(t *testing.T) {
		// Given
		tracker := ctests_tracker.NewCtestsTracker()
		tracker.InsertPackageUt("somePackage 1")
		tracker.InsertPackageUt("somePackage 2")
		tracker.InsertPackageUt("somePackage 3")

		// When
		somePackage3 := tracker.FindPackageWithName("somePackage 3")
		tracker.DeletePackage(somePackage3)

		// Then
		assert.Equal(
			tracker.Packages()[0].Name(),
			"somePackage 1",
		)
		assert.Equal(
			tracker.Packages()[1].Name(),
			"somePackage 2",
		)
		assert.Equal(tracker.PackagesCount(), 2)
	}, t)
}
